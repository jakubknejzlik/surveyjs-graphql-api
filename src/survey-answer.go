package src

import (
	"context"
	"encoding/json"

	"github.com/jakubknejzlik/surveyjs-graphql-api/gen"
)

func getSurveyAnswerValues(ctx context.Context, answer *gen.SurveyAnswer, choicesMap SurveyChoicesMap) (row *gen.SurveyExportRow, err error) {
	answers := map[string]*string{}
	values := []*gen.SurveyExportValue{}
	if answer.Content != nil {
		err = json.Unmarshal([]byte(*answer.Content), &answers)
		if err != nil {
			return
		}

		for key, value := range answers {
			var text string
			if value != nil && choicesMap[key] != nil {
				text = choicesMap[key][*value]
			}
			values = append(values, &gen.SurveyExportValue{
				Key:   key,
				Value: value,
				Text:  &text,
			})
		}
	}
	row = &gen.SurveyExportRow{
		Answer: answer,
		Values: values,
	}
	return
}
