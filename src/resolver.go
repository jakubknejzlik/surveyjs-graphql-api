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
	surveyFields := map[string][]*gen.SurveyExportField{}
	surveyRows := map[string][]*gen.SurveyExportRow{}

	for _, answer := range answers {
		survey, surveyLoaded := surveyMap[*answer.SurveyID]

		if !surveyLoaded {
			survey, err = r.Handlers.SurveyAnswerSurvey(ctx, r.GeneratedResolver, answer)
			if err != nil {
				return
			}
			surveyMap[*answer.SurveyID] = survey
			surveyRows[*answer.SurveyID] = []*gen.SurveyExportRow{}
		}

		_fields, _err := getSurveyFields(ctx, survey)
		if _err != nil {
			err = _err
			return
		}
		if !surveyLoaded {
			for _, field := range _fields {
				surveyFields[*answer.SurveyID] = append(surveyFields[*answer.SurveyID], field)
			}
		}

		row, _err := getSurveyAnswerValues(ctx, answer)
		if _err != nil {
			err = _err
			return
		}
		surveyRows[*answer.SurveyID] = append(surveyRows[*answer.SurveyID], row)
		// rows = append(rows, row)
	}

	items := []*gen.SurveyExportItem{}
	for surveyID, survey := range surveyMap {
		items = append(items, &gen.SurveyExportItem{
			Survey: survey,
			Fields: surveyFields[surveyID],
			Rows:   surveyRows[surveyID],
		})
	}

	// item := &gen.SurveyExportItem{Rows: rows, Fields: fields}
	export = &gen.SurveyExport{Items: items}
	return
}
