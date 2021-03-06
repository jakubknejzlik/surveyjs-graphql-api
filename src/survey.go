package src

import (
	"context"
	"encoding/json"

	"github.com/jakubknejzlik/surveyjs-graphql-api/gen"
)

type SurveyChoicesMap = map[string]map[string]string

// type SurveyPageElementChoice struct {
// 	Value string
// 	Text string
// }
type SurveyPageElement struct {
	Type string
	Name string
	// Title *string
	// Choices []SurveyPageElementChoice
}
type SurveyPage struct {
	Elements []SurveyPageElement
}
type SurveyContent struct {
	Locale string
	Pages  []SurveyPage
}

func getSurveyFields(ctx context.Context, survey *gen.Survey) (fields []*gen.SurveyExportField, err error) {
	fields = []*gen.SurveyExportField{}
	// choicesMap = map[string]map[string]string{}
	if survey != nil && survey.Content != nil {
		content := SurveyContent{}
		err = json.Unmarshal([]byte(*survey.Content), &content)
		if err != nil {
			return
		}

		for _, page := range content.Pages {
			for _, element := range page.Elements {
				fields = append(fields, &gen.SurveyExportField{
					Key:   element.Name,
					Title: nil,
				})
				// for _, choice := range element.Choices {
				// if _, exists := choicesMap[element.Name]; !exists {
				// 	choicesMap[element.Name] = map[string]string{}
				// }
				// choicesMap[element.Name][choice.Value] = choice.Value
				// }
			}
		}
	}
	return
}
