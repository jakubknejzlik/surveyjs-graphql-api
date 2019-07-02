package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/novacloudcz/graphql-orm/events"
	"github.com/novacloudcz/graphql-orm/resolvers"
	uuid "github.com/satori/go.uuid"
)

func getPrincipalID(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
}

type GeneratedResolver struct {
	DB              *DB
	EventController *events.EventController
}

func (r *GeneratedResolver) Mutation() MutationResolver {
	return &GeneratedMutationResolver{r}
}
func (r *GeneratedResolver) Query() QueryResolver {
	return &GeneratedQueryResolver{r}
}

func (r *GeneratedResolver) SurveyResultType() SurveyResultTypeResolver {
	return &GeneratedSurveyResultTypeResolver{r}
}

func (r *GeneratedResolver) Survey() SurveyResolver {
	return &GeneratedSurveyResolver{r}
}

func (r *GeneratedResolver) AnswerResultType() AnswerResultTypeResolver {
	return &GeneratedAnswerResultTypeResolver{r}
}

func (r *GeneratedResolver) Answer() AnswerResolver {
	return &GeneratedAnswerResolver{r}
}

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateSurvey(ctx context.Context, input map[string]interface{}) (item *Survey, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &Survey{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Survey",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		item.Content = changes.Content
		event.AddNewValue("content", changes.Content)
	}

	if ids, ok := input["answersIds"].([]interface{}); ok {
		items := []Answer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateSurvey(ctx context.Context, id string, input map[string]interface{}) (item *Survey, err error) {
	principalID := getPrincipalID(ctx)
	item = &Survey{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Survey",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		event.AddOldValue("content", item.Content)
		event.AddNewValue("content", changes.Content)
		item.Content = changes.Content
	}

	if ids, ok := input["answersIds"].([]interface{}); ok {
		items := []Answer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		data, _ := json.Marshal(event)
		fmt.Println("??", string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteSurvey(ctx context.Context, id string) (item *Survey, err error) {
	item = &Survey{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

func (r *GeneratedMutationResolver) CreateAnswer(ctx context.Context, input map[string]interface{}) (item *Answer, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &Answer{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Answer",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AnswerChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["userID"]; ok && (item.UserID != changes.UserID) {
		item.UserID = changes.UserID
		event.AddNewValue("userID", changes.UserID)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed
		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		item.Content = changes.Content
		event.AddNewValue("content", changes.Content)
	}

	if _, ok := input["surveyId"]; ok && (item.SurveyID != changes.SurveyID) && (item.SurveyID == nil || changes.SurveyID == nil || *item.SurveyID != *changes.SurveyID) {
		item.SurveyID = changes.SurveyID
		event.AddNewValue("surveyId", changes.SurveyID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateAnswer(ctx context.Context, id string, input map[string]interface{}) (item *Answer, err error) {
	principalID := getPrincipalID(ctx)
	item = &Answer{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Answer",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AnswerChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["userID"]; ok && (item.UserID != changes.UserID) {
		event.AddOldValue("userID", item.UserID)
		event.AddNewValue("userID", changes.UserID)
		item.UserID = changes.UserID
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		event.AddOldValue("content", item.Content)
		event.AddNewValue("content", changes.Content)
		item.Content = changes.Content
	}

	if _, ok := input["surveyId"]; ok && (item.SurveyID != changes.SurveyID) && (item.SurveyID == nil || changes.SurveyID == nil || *item.SurveyID != *changes.SurveyID) {
		event.AddOldValue("surveyId", item.SurveyID)
		event.AddNewValue("surveyId", changes.SurveyID)
		item.SurveyID = changes.SurveyID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		data, _ := json.Marshal(event)
		fmt.Println("??", string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAnswer(ctx context.Context, id string) (item *Answer, err error) {
	item = &Answer{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "id = ?", id).Error

	return
}

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) Survey(ctx context.Context, id *string, q *string) (*Survey, error) {
	t := Survey{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Surveys(ctx context.Context, offset *int, limit *int, q *string, sort []SurveySortType, filter *SurveyFilterType) (*SurveyResultType, error) {
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

type GeneratedSurveyResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyResultTypeResolver) Items(ctx context.Context, obj *SurveyResultType) (items []*Survey, err error) {
	err = obj.GetItems(ctx, r.DB.db, "surveys", &items)
	return
}

func (r *GeneratedSurveyResultTypeResolver) Count(ctx context.Context, obj *SurveyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Survey{})
}

type GeneratedSurveyResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyResolver) Answers(ctx context.Context, obj *Survey) (res []*Answer, err error) {

	items := []*Answer{}
	err = r.DB.Query().Model(obj).Related(&items, "Answers").Error
	res = items

	return
}

func (r *GeneratedQueryResolver) Answer(ctx context.Context, id *string, q *string) (*Answer, error) {
	t := Answer{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Answers(ctx context.Context, offset *int, limit *int, q *string, sort []AnswerSortType, filter *AnswerFilterType) (*AnswerResultType, error) {
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

type GeneratedAnswerResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedAnswerResultTypeResolver) Items(ctx context.Context, obj *AnswerResultType) (items []*Answer, err error) {
	err = obj.GetItems(ctx, r.DB.db, "answers", &items)
	return
}

func (r *GeneratedAnswerResultTypeResolver) Count(ctx context.Context, obj *AnswerResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Answer{})
}

type GeneratedAnswerResolver struct{ *GeneratedResolver }

func (r *GeneratedAnswerResolver) Survey(ctx context.Context, obj *Answer) (res *Survey, err error) {

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
