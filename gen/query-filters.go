package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type SurveyQueryFilter struct {
	Query *string
}

func (qf *SurveyQueryFilter) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
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
		if err := qf.applyQueryWithFields(fields, part, "surveys", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *SurveyQueryFilter) applyQueryWithFields(fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}
	aliasPrefix := alias + "."

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["name"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]sname LIKE ? OR %[1]sname LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["content"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]scontent LIKE ? OR %[1]scontent LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_answers"
		*joins = append(*joins, "LEFT JOIN answers "+_alias+" ON "+_alias+".surveyId = "+alias+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := AnswerQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type AnswerQueryFilter struct {
	Query *string
}

func (qf *AnswerQueryFilter) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
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
		if err := qf.applyQueryWithFields(fields, part, "answers", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *AnswerQueryFilter) applyQueryWithFields(fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}
	aliasPrefix := alias + "."

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["content"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]scontent LIKE ? OR %[1]scontent LIKE ?", aliasPrefix))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["employees"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_survey"
		*joins = append(*joins, "LEFT JOIN surveys "+_alias+" ON "+_alias+".id = "+alias+".surveyId")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := SurveyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(_fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
