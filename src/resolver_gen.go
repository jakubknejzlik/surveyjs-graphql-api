package src

import (
	"context"

	"github.com/jakubknejzlik/surveyjs-graphql-api/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := gen.GetTransaction(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type SurveyResultTypeResolver struct {
	*gen.GeneratedSurveyResultTypeResolver
}

func (r *Resolver) SurveyResultType() gen.SurveyResultTypeResolver {
	return &SurveyResultTypeResolver{&gen.GeneratedSurveyResultTypeResolver{r.GeneratedResolver}}
}

type SurveyResolver struct {
	*gen.GeneratedSurveyResolver
}

func (r *Resolver) Survey() gen.SurveyResolver {
	return &SurveyResolver{&gen.GeneratedSurveyResolver{r.GeneratedResolver}}
}

type SurveyAnswerResultTypeResolver struct {
	*gen.GeneratedSurveyAnswerResultTypeResolver
}

func (r *Resolver) SurveyAnswerResultType() gen.SurveyAnswerResultTypeResolver {
	return &SurveyAnswerResultTypeResolver{&gen.GeneratedSurveyAnswerResultTypeResolver{r.GeneratedResolver}}
}

type SurveyAnswerResolver struct {
	*gen.GeneratedSurveyAnswerResolver
}

func (r *Resolver) SurveyAnswer() gen.SurveyAnswerResolver {
	return &SurveyAnswerResolver{&gen.GeneratedSurveyAnswerResolver{r.GeneratedResolver}}
}
