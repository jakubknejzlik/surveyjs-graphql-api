package src

import (
	"context"
	"encoding/json"

	"github.com/jakubknejzlik/surveyjs-graphql-api/gen"
)

func getSurveyAnswerValues(ctx context.Context, answer *gen.SurveyAnswer) (row *gen.SurveyExportRow, err error) {
	answers := map[string]*json.RawMessage{}
	values := []*gen.SurveyExportValue{}
	if answer.Content != nil {
		err = json.Unmarshal([]byte(*answer.Content), &answers)
		if err != nil {
			return
		}

		for key, value := range answers {
			_value := ""
			err := json.Unmarshal(*value, &_value)
			if err != nil {
				values := []string{}
				json.Unmarshal(*value, &values)
				_value = values[0]
			}
			values = append(values, &gen.SurveyExportValue{
				Key:   key,
				Value: &_value,
			})
		}
	}
	row = &gen.SurveyExportRow{
		Answer: answer,
		Values: values,
	}
	return
}
