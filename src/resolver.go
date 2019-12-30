package src

import (
	"context"

	"github.com/jakubknejzlik/surveyjs-graphql-api/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateCompany = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return nil, fmt.Errorf("can't touch this!")
	// }

	return resolver
}

func (r *QueryResolver) SurveyExport(ctx context.Context, filter *gen.SurveyExportFilterType) (export *gen.SurveyExport, err error) {
	answer, err := r.Handlers.QuerySurveyAnswer(ctx, r.GeneratedResolver, gen.QuerySurveyAnswerHandlerOptions{
		ID: &filter.AnswerIDs[0],
	})
	if err != nil {
		return
	}

	survey, err := r.Handlers.SurveyAnswerSurvey(ctx, r.GeneratedResolver, answer)
	if err != nil {
		return
	}

	fields, choicesMap, err := getSurveyFields(ctx, survey)
	if err != nil {
		return
	}

	rows := []*gen.SurveyExportRow{}
	row, err := getSurveyAnswerValues(ctx, answer, choicesMap)
	if err != nil {
		return
	}
	rows = append(rows, row)

	export = &gen.SurveyExport{Rows: rows, Fields: fields}
	return
}
