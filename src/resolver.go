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

	answersResult, err := r.Handlers.QuerySurveyAnswers(ctx, r.GeneratedResolver, gen.QuerySurveyAnswersHandlerOptions{
		Filter: &gen.SurveyAnswerFilterType{
			IDIn: filter.AnswerIDs,
		},
	})
	if err != nil {
		return
	}
	var answers []*gen.SurveyAnswer
	err = answersResult.GetItems(ctx, r.DB.Query(), gen.GetItemsOptions{}, &answers)
	if err != nil {
		return
	}

	surveyMap := map[string]*gen.Survey{}
	rows := []*gen.SurveyExportRow{}

	fields := []*gen.SurveyExportField{}
	for _, answer := range answers {
		survey, surveyLoaded := surveyMap[*answer.SurveyID]

		if !surveyLoaded {
			survey, err = r.Handlers.SurveyAnswerSurvey(ctx, r.GeneratedResolver, answer)
			if err != nil {
				return
			}
			surveyMap[*answer.SurveyID] = survey
		}

		_fields, choicesMap, _err := getSurveyFields(ctx, survey)
		if _err != nil {
			err = _err
			return
		}
		if !surveyLoaded {
			for _, field := range _fields {
				fields = append(fields, field)
			}
		}

		row, _err := getSurveyAnswerValues(ctx, answer, choicesMap)
		if _err != nil {
			err = _err
			return
		}
		rows = append(rows, row)
	}

	export = &gen.SurveyExport{Rows: rows, Fields: fields}
	return
}
