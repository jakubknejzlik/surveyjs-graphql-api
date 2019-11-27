package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateSurvey     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Survey, err error)
	UpdateSurvey     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Survey, err error)
	DeleteSurvey     func(ctx context.Context, r *GeneratedResolver, id string) (item *Survey, err error)
	DeleteAllSurveys func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QuerySurvey      func(ctx context.Context, r *GeneratedResolver, opts QuerySurveyHandlerOptions) (*Survey, error)
	QuerySurveys     func(ctx context.Context, r *GeneratedResolver, opts QuerySurveysHandlerOptions) (*SurveyResultType, error)

	SurveyAnswers func(ctx context.Context, r *GeneratedSurveyResolver, obj *Survey) (res []*SurveyAnswer, err error)

	CreateSurveyAnswer     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *SurveyAnswer, err error)
	UpdateSurveyAnswer     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *SurveyAnswer, err error)
	DeleteSurveyAnswer     func(ctx context.Context, r *GeneratedResolver, id string) (item *SurveyAnswer, err error)
	DeleteAllSurveyAnswers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QuerySurveyAnswer      func(ctx context.Context, r *GeneratedResolver, opts QuerySurveyAnswerHandlerOptions) (*SurveyAnswer, error)
	QuerySurveyAnswers     func(ctx context.Context, r *GeneratedResolver, opts QuerySurveyAnswersHandlerOptions) (*SurveyAnswerResultType, error)

	SurveyAnswerSurvey func(ctx context.Context, r *GeneratedSurveyAnswerResolver, obj *SurveyAnswer) (res *Survey, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateSurvey:     CreateSurveyHandler,
		UpdateSurvey:     UpdateSurveyHandler,
		DeleteSurvey:     DeleteSurveyHandler,
		DeleteAllSurveys: DeleteAllSurveysHandler,
		QuerySurvey:      QuerySurveyHandler,
		QuerySurveys:     QuerySurveysHandler,

		SurveyAnswers: SurveyAnswersHandler,

		CreateSurveyAnswer:     CreateSurveyAnswerHandler,
		UpdateSurveyAnswer:     UpdateSurveyAnswerHandler,
		DeleteSurveyAnswer:     DeleteSurveyAnswerHandler,
		DeleteAllSurveyAnswers: DeleteAllSurveyAnswersHandler,
		QuerySurveyAnswer:      QuerySurveyAnswerHandler,
		QuerySurveyAnswers:     QuerySurveyAnswersHandler,

		SurveyAnswerSurvey: SurveyAnswerSurveyHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
