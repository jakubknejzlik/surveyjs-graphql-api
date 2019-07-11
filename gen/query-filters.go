package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type SurveyQueryFilter struct {
	Query *string
}

func (qf *SurveyQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(dialect, fields, part, "surveys", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *SurveyQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["name"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("name")+" LIKE ? OR %[1]s"+dialect.Quote("name")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["content"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("content")+" LIKE ? OR %[1]s"+dialect.Quote("content")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["answers"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_answers"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("survey_answers")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("surveyId")+" = "+dialect.Quote(alias)+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := SurveyAnswerQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type SurveyAnswerQueryFilter struct {
	Query *string
}

func (qf *SurveyAnswerQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(dialect, fields, part, "survey_answers", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *SurveyAnswerQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["content"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("content")+" LIKE ? OR %[1]s"+dialect.Quote("content")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["survey"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_survey"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("surveys")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("surveyId"))

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := SurveyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
