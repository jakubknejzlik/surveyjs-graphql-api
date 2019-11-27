package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QuerySurveyHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *SurveyFilterType
}

func (r *GeneratedQueryResolver) Survey(ctx context.Context, id *string, q *string, filter *SurveyFilterType) (*Survey, error) {
	opts := QuerySurveyHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QuerySurvey(ctx, r.GeneratedResolver, opts)
}
func QuerySurveyHandler(ctx context.Context, r *GeneratedResolver, opts QuerySurveyHandlerOptions) (*Survey, error) {
	query := SurveyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &SurveyResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("surveys")+".id = ?", *opts.ID)
	}

	var items []*Survey
	giOpts := GetItemsOptions{
		Alias:      TableName("surveys"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Survey"}
	}
	return items[0], err
}

type QuerySurveysHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*SurveySortType
	Filter *SurveyFilterType
}

func (r *GeneratedQueryResolver) Surveys(ctx context.Context, offset *int, limit *int, q *string, sort []*SurveySortType, filter *SurveyFilterType) (*SurveyResultType, error) {
	opts := QuerySurveysHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QuerySurveys(ctx, r.GeneratedResolver, opts)
}
func QuerySurveysHandler(ctx context.Context, r *GeneratedResolver, opts QuerySurveysHandlerOptions) (*SurveyResultType, error) {
	query := SurveyQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &SurveyResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedSurveyResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyResultTypeResolver) Items(ctx context.Context, obj *SurveyResultType) (items []*Survey, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("surveys"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedSurveyResultTypeResolver) Count(ctx context.Context, obj *SurveyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Survey{})
}

type GeneratedSurveyResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyResolver) Answers(ctx context.Context, obj *Survey) (res []*SurveyAnswer, err error) {
	return r.Handlers.SurveyAnswers(ctx, r, obj)
}
func SurveyAnswersHandler(ctx context.Context, r *GeneratedSurveyResolver, obj *Survey) (res []*SurveyAnswer, err error) {

	items := []*SurveyAnswer{}
	err = r.DB.Query().Model(obj).Related(&items, "Answers").Error
	res = items

	return
}

func (r *GeneratedSurveyResolver) AnswersIds(ctx context.Context, obj *Survey) (ids []string, err error) {
	ids = []string{}

	items := []*SurveyAnswer{}
	err = r.DB.Query().Model(obj).Select(TableName("survey_answers")+".id").Related(&items, "Answers").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QuerySurveyAnswerHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *SurveyAnswerFilterType
}

func (r *GeneratedQueryResolver) SurveyAnswer(ctx context.Context, id *string, q *string, filter *SurveyAnswerFilterType) (*SurveyAnswer, error) {
	opts := QuerySurveyAnswerHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QuerySurveyAnswer(ctx, r.GeneratedResolver, opts)
}
func QuerySurveyAnswerHandler(ctx context.Context, r *GeneratedResolver, opts QuerySurveyAnswerHandlerOptions) (*SurveyAnswer, error) {
	query := SurveyAnswerQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &SurveyAnswerResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("survey_answers")+".id = ?", *opts.ID)
	}

	var items []*SurveyAnswer
	giOpts := GetItemsOptions{
		Alias:      TableName("survey_answers"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "SurveyAnswer"}
	}
	return items[0], err
}

type QuerySurveyAnswersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*SurveyAnswerSortType
	Filter *SurveyAnswerFilterType
}

func (r *GeneratedQueryResolver) SurveyAnswers(ctx context.Context, offset *int, limit *int, q *string, sort []*SurveyAnswerSortType, filter *SurveyAnswerFilterType) (*SurveyAnswerResultType, error) {
	opts := QuerySurveyAnswersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QuerySurveyAnswers(ctx, r.GeneratedResolver, opts)
}
func QuerySurveyAnswersHandler(ctx context.Context, r *GeneratedResolver, opts QuerySurveyAnswersHandlerOptions) (*SurveyAnswerResultType, error) {
	query := SurveyAnswerQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &SurveyAnswerResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedSurveyAnswerResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyAnswerResultTypeResolver) Items(ctx context.Context, obj *SurveyAnswerResultType) (items []*SurveyAnswer, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("survey_answers"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedSurveyAnswerResultTypeResolver) Count(ctx context.Context, obj *SurveyAnswerResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &SurveyAnswer{})
}

type GeneratedSurveyAnswerResolver struct{ *GeneratedResolver }

func (r *GeneratedSurveyAnswerResolver) Survey(ctx context.Context, obj *SurveyAnswer) (res *Survey, err error) {
	return r.Handlers.SurveyAnswerSurvey(ctx, r, obj)
}
func SurveyAnswerSurveyHandler(ctx context.Context, r *GeneratedSurveyAnswerResolver, obj *SurveyAnswer) (res *Survey, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.SurveyID != nil {
		item, _err := loaders["Survey"].Load(ctx, dataloader.StringKey(*obj.SurveyID))()
		res, _ = item.(*Survey)

		err = _err
	}

	return
}
