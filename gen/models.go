package gen

import (
	"github.com/novacloudcz/graphql-orm/resolvers"
	"time"
)

type SurveyResultType struct {
	resolvers.EntityResultType
}

type Survey struct {
	ID        string    `json:"id" gorm:"column:id;primary_key"`
	Name      *string   `json:"name" gorm:"column:name"`
	Content   *string   `json:"content" gorm:"column:content"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`

	Answers []*Answer `json:"answers" gorm:"foreignkey:SurveyID"`
}

type AnswerResultType struct {
	resolvers.EntityResultType
}

type Answer struct {
	ID        string    `json:"id" gorm:"column:id;primary_key"`
	UserID    string    `json:"userID" gorm:"column:userID"`
	Completed *bool     `json:"completed" gorm:"column:completed"`
	Content   *string   `json:"content" gorm:"column:content"`
	SurveyID  *string   `json:"surveyId" gorm:"column:surveyId"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`

	Survey *Survey `json:"survey"`
}
