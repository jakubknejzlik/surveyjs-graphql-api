package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

type Query {
  survey(id: ID, q: String, filter: SurveyFilterType): Survey
  surveys(offset: Int, limit: Int = 30, q: String, sort: [SurveySortType!], filter: SurveyFilterType): SurveyResultType
  surveyAnswer(id: ID, q: String, filter: SurveyAnswerFilterType): SurveyAnswer
  surveyAnswers(offset: Int, limit: Int = 30, q: String, sort: [SurveyAnswerSortType!], filter: SurveyAnswerFilterType): SurveyAnswerResultType
}

type Mutation {
  createSurvey(input: SurveyCreateInput!): Survey!
  updateSurvey(id: ID!, input: SurveyUpdateInput!): Survey!
  deleteSurvey(id: ID!): Survey!
  deleteAllSurveys: Boolean!
  createSurveyAnswer(input: SurveyAnswerCreateInput!): SurveyAnswer!
  updateSurveyAnswer(id: ID!, input: SurveyAnswerUpdateInput!): SurveyAnswer!
  deleteSurveyAnswer(id: ID!): SurveyAnswer!
  deleteAllSurveyAnswers: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

type Survey @key(fields: "id") {
  id: ID!
  name: String
  content: String
  answers: [SurveyAnswer!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  answersIds: [ID!]!
}

type SurveyAnswer @key(fields: "id") {
  id: ID!
  completed: Boolean
  content: String
  survey: Survey
  surveyId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

union _Entity = Survey | SurveyAnswer

input SurveyCreateInput {
  id: ID
  name: String
  content: String
  answersIds: [ID!]
}

input SurveyUpdateInput {
  name: String
  content: String
  answersIds: [ID!]
}

input SurveySortType {
  id: ObjectSortType
  name: ObjectSortType
  content: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  answersIds: ObjectSortType
  answers: SurveyAnswerSortType
}

input SurveyFilterType {
  AND: [SurveyFilterType!]
  OR: [SurveyFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
  content: String
  content_ne: String
  content_gt: String
  content_lt: String
  content_gte: String
  content_lte: String
  content_in: [String!]
  content_like: String
  content_prefix: String
  content_suffix: String
  content_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  answers: SurveyAnswerFilterType
}

type SurveyResultType {
  items: [Survey!]!
  count: Int!
}

input SurveyAnswerCreateInput {
  id: ID
  completed: Boolean
  content: String
  surveyId: ID
}

input SurveyAnswerUpdateInput {
  completed: Boolean
  content: String
  surveyId: ID
}

input SurveyAnswerSortType {
  id: ObjectSortType
  completed: ObjectSortType
  content: ObjectSortType
  surveyId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  survey: SurveySortType
}

input SurveyAnswerFilterType {
  AND: [SurveyAnswerFilterType!]
  OR: [SurveyAnswerFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  completed: Boolean
  completed_ne: Boolean
  completed_gt: Boolean
  completed_lt: Boolean
  completed_gte: Boolean
  completed_lte: Boolean
  completed_in: [Boolean!]
  completed_null: Boolean
  content: String
  content_ne: String
  content_gt: String
  content_lt: String
  content_gte: String
  content_lte: String
  content_in: [String!]
  content_like: String
  content_prefix: String
  content_suffix: String
  content_null: Boolean
  surveyId: ID
  surveyId_ne: ID
  surveyId_gt: ID
  surveyId_lt: ID
  surveyId_gte: ID
  surveyId_lte: ID
  surveyId_in: [ID!]
  surveyId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  survey: SurveyFilterType
}

type SurveyAnswerResultType {
  items: [SurveyAnswer!]!
  count: Int!
}`
)
