package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *SurveyFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "surveys", wheres, values, joins)
}
func (f *SurveyFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Answers != nil {
		_alias := alias + "_answers"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("survey_answers")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("surveyId")+" = "+dialect.Quote(alias)+".id")
		err := f.Answers.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *SurveyFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.Content != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" = ?")
		values = append(values, f.Content)
	}
	if f.ContentNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" != ?")
		values = append(values, f.ContentNe)
	}
	if f.ContentGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" > ?")
		values = append(values, f.ContentGt)
	}
	if f.ContentLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" < ?")
		values = append(values, f.ContentLt)
	}
	if f.ContentGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" >= ?")
		values = append(values, f.ContentGte)
	}
	if f.ContentLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" <= ?")
		values = append(values, f.ContentLte)
	}
	if f.ContentIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" IN (?)")
		values = append(values, f.ContentIn)
	}
	if f.ContentLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ContentLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ContentPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentPrefix))
	}
	if f.ContentSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *SurveyFilterType) AndWith(f2 ...*SurveyFilterType) *SurveyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &SurveyFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *SurveyFilterType) OrWith(f2 ...*SurveyFilterType) *SurveyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &SurveyFilterType{
		Or: append(_f2, f),
	}
}

func (f *SurveyAnswerFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "survey_answers", wheres, values, joins)
}
func (f *SurveyAnswerFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Survey != nil {
		_alias := alias + "_survey"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("surveys")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("surveyId"))
		err := f.Survey.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *SurveyAnswerFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" = ?")
		values = append(values, f.Completed)
	}
	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}
	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}
	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}
	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}
	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}
	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.Content != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" = ?")
		values = append(values, f.Content)
	}
	if f.ContentNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" != ?")
		values = append(values, f.ContentNe)
	}
	if f.ContentGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" > ?")
		values = append(values, f.ContentGt)
	}
	if f.ContentLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" < ?")
		values = append(values, f.ContentLt)
	}
	if f.ContentGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" >= ?")
		values = append(values, f.ContentGte)
	}
	if f.ContentLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" <= ?")
		values = append(values, f.ContentLte)
	}
	if f.ContentIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" IN (?)")
		values = append(values, f.ContentIn)
	}
	if f.ContentLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ContentLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ContentPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentPrefix))
	}
	if f.ContentSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("content")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentSuffix))
	}

	if f.SurveyID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" = ?")
		values = append(values, f.SurveyID)
	}
	if f.SurveyIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" != ?")
		values = append(values, f.SurveyIDNe)
	}
	if f.SurveyIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" > ?")
		values = append(values, f.SurveyIDGt)
	}
	if f.SurveyIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" < ?")
		values = append(values, f.SurveyIDLt)
	}
	if f.SurveyIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" >= ?")
		values = append(values, f.SurveyIDGte)
	}
	if f.SurveyIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" <= ?")
		values = append(values, f.SurveyIDLte)
	}
	if f.SurveyIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("surveyId")+" IN (?)")
		values = append(values, f.SurveyIDIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *SurveyAnswerFilterType) AndWith(f2 ...*SurveyAnswerFilterType) *SurveyAnswerFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &SurveyAnswerFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *SurveyAnswerFilterType) OrWith(f2 ...*SurveyAnswerFilterType) *SurveyAnswerFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &SurveyAnswerFilterType{
		Or: append(_f2, f),
	}
}
