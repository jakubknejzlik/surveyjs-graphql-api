package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := GetTransaction(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func GetTransaction(ctx context.Context) *gorm.DB {
	return ctx.Value(KeyMutationTransaction).(*gorm.DB)
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateSurvey(ctx context.Context, input map[string]interface{}) (item *Survey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateSurvey(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateSurveyHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Survey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Survey{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Survey",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		item.Content = changes.Content

		event.AddNewValue("content", changes.Content)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["answersIds"]; exists {
		items := []SurveyAnswer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateSurvey(ctx context.Context, id string, input map[string]interface{}) (item *Survey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateSurvey(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateSurveyHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Survey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Survey{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Survey",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		event.AddOldValue("content", item.Content)
		event.AddNewValue("content", changes.Content)
		item.Content = changes.Content
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["answersIds"]; exists {
		items := []SurveyAnswer{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Answers")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteSurvey(ctx context.Context, id string) (item *Survey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteSurvey(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteSurveyHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Survey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Survey{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Survey",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("surveys")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllSurveys(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllSurveys(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllSurveysHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&Survey{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateSurveyAnswer(ctx context.Context, input map[string]interface{}) (item *SurveyAnswer, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateSurveyAnswer(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateSurveyAnswerHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *SurveyAnswer, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &SurveyAnswer{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "SurveyAnswer",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyAnswerChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed

		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		item.Content = changes.Content

		event.AddNewValue("content", changes.Content)
	}

	if _, ok := input["surveyId"]; ok && (item.SurveyID != changes.SurveyID) && (item.SurveyID == nil || changes.SurveyID == nil || *item.SurveyID != *changes.SurveyID) {
		item.SurveyID = changes.SurveyID

		event.AddNewValue("surveyId", changes.SurveyID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateSurveyAnswer(ctx context.Context, id string, input map[string]interface{}) (item *SurveyAnswer, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateSurveyAnswer(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateSurveyAnswerHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *SurveyAnswer, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &SurveyAnswer{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "SurveyAnswer",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes SurveyAnswerChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["content"]; ok && (item.Content != changes.Content) && (item.Content == nil || changes.Content == nil || *item.Content != *changes.Content) {
		event.AddOldValue("content", item.Content)
		event.AddNewValue("content", changes.Content)
		item.Content = changes.Content
	}

	if _, ok := input["surveyId"]; ok && (item.SurveyID != changes.SurveyID) && (item.SurveyID == nil || changes.SurveyID == nil || *item.SurveyID != *changes.SurveyID) {
		event.AddOldValue("surveyId", item.SurveyID)
		event.AddNewValue("surveyId", changes.SurveyID)
		item.SurveyID = changes.SurveyID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteSurveyAnswer(ctx context.Context, id string) (item *SurveyAnswer, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteSurveyAnswer(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteSurveyAnswerHandler(ctx context.Context, r *GeneratedResolver, id string) (item *SurveyAnswer, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &SurveyAnswer{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "SurveyAnswer",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("survey_answers")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllSurveyAnswers(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllSurveyAnswers(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllSurveyAnswersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&SurveyAnswer{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
