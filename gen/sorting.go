package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s SurveySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("surveys"), sorts, joins)
}
func (s SurveySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+"name "+s.Name.String())
	}

	if s.Content != nil {
		*sorts = append(*sorts, aliasPrefix+"content "+s.Content.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Answers != nil {
		_alias := alias + "_answers"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("survey_answers"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("surveyId")+" = "+dialect.Quote(alias)+".id")
		err := s.Answers.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s SurveyAnswerSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("survey_answers"), sorts, joins)
}
func (s SurveyAnswerSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Completed != nil {
		*sorts = append(*sorts, aliasPrefix+"completed "+s.Completed.String())
	}

	if s.Content != nil {
		*sorts = append(*sorts, aliasPrefix+"content "+s.Content.String())
	}

	if s.SurveyID != nil {
		*sorts = append(*sorts, aliasPrefix+"surveyId "+s.SurveyID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Survey != nil {
		_alias := alias + "_survey"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("surveys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("surveyId"))
		err := s.Survey.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
