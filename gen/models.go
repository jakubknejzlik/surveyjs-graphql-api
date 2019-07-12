package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/novacloudcz/graphql-orm/resolvers"
)

type SurveyResultType struct {
	resolvers.EntityResultType
}

type Survey struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	Content   *string    `json:"content" gorm:"column:content;type:text"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Answers []*SurveyAnswer `json:"answers" gorm:"foreignkey:SurveyID"`
}

type SurveyChanges struct {
	ID        string
	Name      *string
	Content   *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string
}

type SurveyAnswerResultType struct {
	resolvers.EntityResultType
}

type SurveyAnswer struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Completed *bool      `json:"completed" gorm:"column:completed"`
	Content   *string    `json:"content" gorm:"column:content;type:text"`
	SurveyID  *string    `json:"surveyId" gorm:"column:surveyId"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Survey *Survey `json:"survey"`
}

type SurveyAnswerChanges struct {
	ID        string
	Completed *bool
	Content   *string
	SurveyID  *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
