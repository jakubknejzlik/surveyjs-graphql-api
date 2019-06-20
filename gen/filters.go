package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *SurveyFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, "surveys", wheres, values, joins)
}
func (f *SurveyFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Answers != nil {
		_alias := alias + "_answers"
		*joins = append(*joins, "LEFT JOIN answers "+_alias+" ON "+_alias+".surveyId = "+alias+".id")
		err := f.Answers.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *SurveyFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, "("+strings.Join(cs, " OR ")+")")
		values = append(values, vs...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, strings.Join(cs, " AND "))
		values = append(values, vs...)
	}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+"id = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+"id != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+"id > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+"id < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+"id >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+"id <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+"id IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+"name = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+"name != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+"name > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+"name < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+"name >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+"name <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+"name IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.NameLike, "?", "_"), "*", "%"))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+"name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.Content != nil {
		conditions = append(conditions, aliasPrefix+"content = ?")
		values = append(values, f.Content)
	}
	if f.ContentNe != nil {
		conditions = append(conditions, aliasPrefix+"content != ?")
		values = append(values, f.ContentNe)
	}
	if f.ContentGt != nil {
		conditions = append(conditions, aliasPrefix+"content > ?")
		values = append(values, f.ContentGt)
	}
	if f.ContentLt != nil {
		conditions = append(conditions, aliasPrefix+"content < ?")
		values = append(values, f.ContentLt)
	}
	if f.ContentGte != nil {
		conditions = append(conditions, aliasPrefix+"content >= ?")
		values = append(values, f.ContentGte)
	}
	if f.ContentLte != nil {
		conditions = append(conditions, aliasPrefix+"content <= ?")
		values = append(values, f.ContentLte)
	}
	if f.ContentIn != nil {
		conditions = append(conditions, aliasPrefix+"content IN (?)")
		values = append(values, f.ContentIn)
	}
	if f.ContentLike != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.ContentLike, "?", "_"), "*", "%"))
	}
	if f.ContentPrefix != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentPrefix))
	}
	if f.ContentSuffix != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentSuffix))
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"createdAt != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"createdAt IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	return
}

func (f *AnswerFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, "answers", wheres, values, joins)
}
func (f *AnswerFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Survey != nil {
		_alias := alias + "_survey"
		*joins = append(*joins, "LEFT JOIN surveys "+_alias+" ON "+_alias+".id = "+alias+".surveyId")
		err := f.Survey.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *AnswerFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, "("+strings.Join(cs, " OR ")+")")
		values = append(values, vs...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		for _, or := range f.Or {
			_cond, _values := or.WhereContent(aliasPrefix)
			cs = append(cs, _cond...)
			vs = append(vs, _values...)
		}
		conditions = append(conditions, strings.Join(cs, " AND "))
		values = append(values, vs...)
	}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+"id = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+"id != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+"id > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+"id < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+"id >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+"id <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+"id IN (?)")
		values = append(values, f.IDIn)
	}

	if f.UserID != nil {
		conditions = append(conditions, aliasPrefix+"userID = ?")
		values = append(values, f.UserID)
	}
	if f.UserIDNe != nil {
		conditions = append(conditions, aliasPrefix+"userID != ?")
		values = append(values, f.UserIDNe)
	}
	if f.UserIDGt != nil {
		conditions = append(conditions, aliasPrefix+"userID > ?")
		values = append(values, f.UserIDGt)
	}
	if f.UserIDLt != nil {
		conditions = append(conditions, aliasPrefix+"userID < ?")
		values = append(values, f.UserIDLt)
	}
	if f.UserIDGte != nil {
		conditions = append(conditions, aliasPrefix+"userID >= ?")
		values = append(values, f.UserIDGte)
	}
	if f.UserIDLte != nil {
		conditions = append(conditions, aliasPrefix+"userID <= ?")
		values = append(values, f.UserIDLte)
	}
	if f.UserIDIn != nil {
		conditions = append(conditions, aliasPrefix+"userID IN (?)")
		values = append(values, f.UserIDIn)
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+"completed = ?")
		values = append(values, f.Completed)
	}
	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+"completed != ?")
		values = append(values, f.CompletedNe)
	}
	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+"completed > ?")
		values = append(values, f.CompletedGt)
	}
	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+"completed < ?")
		values = append(values, f.CompletedLt)
	}
	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+"completed >= ?")
		values = append(values, f.CompletedGte)
	}
	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+"completed <= ?")
		values = append(values, f.CompletedLte)
	}
	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+"completed IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.Content != nil {
		conditions = append(conditions, aliasPrefix+"content = ?")
		values = append(values, f.Content)
	}
	if f.ContentNe != nil {
		conditions = append(conditions, aliasPrefix+"content != ?")
		values = append(values, f.ContentNe)
	}
	if f.ContentGt != nil {
		conditions = append(conditions, aliasPrefix+"content > ?")
		values = append(values, f.ContentGt)
	}
	if f.ContentLt != nil {
		conditions = append(conditions, aliasPrefix+"content < ?")
		values = append(values, f.ContentLt)
	}
	if f.ContentGte != nil {
		conditions = append(conditions, aliasPrefix+"content >= ?")
		values = append(values, f.ContentGte)
	}
	if f.ContentLte != nil {
		conditions = append(conditions, aliasPrefix+"content <= ?")
		values = append(values, f.ContentLte)
	}
	if f.ContentIn != nil {
		conditions = append(conditions, aliasPrefix+"content IN (?)")
		values = append(values, f.ContentIn)
	}
	if f.ContentLike != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, strings.ReplaceAll(strings.ReplaceAll(*f.ContentLike, "?", "_"), "*", "%"))
	}
	if f.ContentPrefix != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentPrefix))
	}
	if f.ContentSuffix != nil {
		conditions = append(conditions, aliasPrefix+"content LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentSuffix))
	}

	if f.SurveyID != nil {
		conditions = append(conditions, aliasPrefix+"surveyId = ?")
		values = append(values, f.SurveyID)
	}
	if f.SurveyIDNe != nil {
		conditions = append(conditions, aliasPrefix+"surveyId != ?")
		values = append(values, f.SurveyIDNe)
	}
	if f.SurveyIDGt != nil {
		conditions = append(conditions, aliasPrefix+"surveyId > ?")
		values = append(values, f.SurveyIDGt)
	}
	if f.SurveyIDLt != nil {
		conditions = append(conditions, aliasPrefix+"surveyId < ?")
		values = append(values, f.SurveyIDLt)
	}
	if f.SurveyIDGte != nil {
		conditions = append(conditions, aliasPrefix+"surveyId >= ?")
		values = append(values, f.SurveyIDGte)
	}
	if f.SurveyIDLte != nil {
		conditions = append(conditions, aliasPrefix+"surveyId <= ?")
		values = append(values, f.SurveyIDLte)
	}
	if f.SurveyIDIn != nil {
		conditions = append(conditions, aliasPrefix+"surveyId IN (?)")
		values = append(values, f.SurveyIDIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"updatedAt IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+"createdAt != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+"createdAt < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+"createdAt <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+"createdAt IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	return
}
