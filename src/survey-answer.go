package src

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
				numVal := -1
				err := json.Unmarshal(*value, &numVal)
				if numVal == -1 {
					_value = "-"
				} else {
					_value = fmt.Sprintf("%d", numVal)
				}
				if err != nil {
					values := []string{}
					err := json.Unmarshal(*value, &values)
					if err != nil {
						return nil, err
					}
					if len(values) > 0 {
						_value = strings.Join(values, ",")
					} else {
						_value = "–"
					}
				}
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
