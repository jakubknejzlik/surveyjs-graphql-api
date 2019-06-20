package gen

import (
	"context"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
)

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

type Resolver struct {
	DB *DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) SurveyResultType() SurveyResultTypeResolver {
	return &surveyResultTypeResolver{r}
}

func (r *Resolver) Survey() SurveyResolver {
	return &surveyResolver{r}
}

func (r *Resolver) AnswerResultType() AnswerResultTypeResolver {
	return &answerResultTypeResolver{r}
}

func (r *Resolver) Answer() AnswerResolver {
	return &answerResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateSurvey(ctx context.Context, input map[string]interface{}) (item *Survey, err error) {
	ID, ok := input["id"].(string)
	if !ok || ID == "" {
		ID = uuid.Must(uuid.NewV4()).String()
	}
	item = &Survey{ID: ID}
	tx := r.DB.db.Begin()

	if ids, ok := input["answersIds"].([]interface{}); ok {
		items := []Answer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateSurvey(ctx context.Context, id string, input map[string]interface{}) (item *Survey, err error) {
	item = &Survey{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	if ids, ok := input["answersIds"].([]interface{}); ok {
		items := []Answer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteSurvey(ctx context.Context, id string) (item *Survey, err error) {
	item = &Survey{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *mutationResolver) CreateAnswer(ctx context.Context, input map[string]interface{}) (item *Answer, err error) {
	ID, ok := input["id"].(string)
	if !ok || ID == "" {
		ID = uuid.Must(uuid.NewV4()).String()
	}
	item = &Answer{ID: ID}
	tx := r.DB.db.Begin()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) UpdateAnswer(ctx context.Context, id string, input map[string]interface{}) (item *Answer, err error) {
	item = &Answer{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: item,
	})
	if err != nil {
		tx.Rollback()
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
func (r *mutationResolver) DeleteAnswer(ctx context.Context, id string) (item *Answer, err error) {
	item = &Answer{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Survey(ctx context.Context, id *string, q *string) (*Survey, error) {
	t := Survey{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Surveys(ctx context.Context, offset *int, limit *int, q *string, sort []SurveySortType, filter *SurveyFilterType) (*SurveyResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := SurveyQueryFilter{q}
	return &SurveyResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type surveyResultTypeResolver struct{ *Resolver }

func (r *surveyResultTypeResolver) Items(ctx context.Context, obj *SurveyResultType) (items []*Survey, err error) {
	err = obj.GetItems(ctx, r.DB.db, "surveys", &items)
	return
}

func (r *surveyResultTypeResolver) Count(ctx context.Context, obj *SurveyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Survey{})
}

type surveyResolver struct{ *Resolver }

func (r *surveyResolver) Answers(ctx context.Context, obj *Survey) (res []*Answer, err error) {

	items := []*Answer{}
	err = r.DB.Query().Model(obj).Related(&items, "Answers").Error
	res = items

	return
}

func (r *queryResolver) Answer(ctx context.Context, id *string, q *string) (*Answer, error) {
	t := Answer{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *queryResolver) Answers(ctx context.Context, offset *int, limit *int, q *string, sort []AnswerSortType, filter *AnswerFilterType) (*AnswerResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := AnswerQueryFilter{q}
	return &AnswerResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type answerResultTypeResolver struct{ *Resolver }

func (r *answerResultTypeResolver) Items(ctx context.Context, obj *AnswerResultType) (items []*Answer, err error) {
	err = obj.GetItems(ctx, r.DB.db, "answers", &items)
	return
}

func (r *answerResultTypeResolver) Count(ctx context.Context, obj *AnswerResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Answer{})
}

type answerResolver struct{ *Resolver }

func (r *answerResolver) Survey(ctx context.Context, obj *Answer) (res *Survey, err error) {

	item := Survey{}
	_res := r.DB.Query().Model(obj).Related(&item, "Survey")
	if _res.RecordNotFound() {
		return
	} else {
		err = _res.Error
	}
	res = &item

	return
}
