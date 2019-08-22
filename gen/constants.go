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

enum SurveySortType {
  ID_ASC
  ID_DESC
  NAME_ASC
  NAME_DESC
  CONTENT_ASC
  CONTENT_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  ANSWERS_IDS_ASC
  ANSWERS_IDS_DESC
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
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
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

enum SurveyAnswerSortType {
  ID_ASC
  ID_DESC
  COMPLETED_ASC
  COMPLETED_DESC
  CONTENT_ASC
  CONTENT_DESC
  SURVEY_ID_ASC
  SURVEY_ID_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
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
  completed: Boolean
  completed_ne: Boolean
  completed_gt: Boolean
  completed_lt: Boolean
  completed_gte: Boolean
  completed_lte: Boolean
  completed_in: [Boolean!]
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
  surveyId: ID
  surveyId_ne: ID
  surveyId_gt: ID
  surveyId_lt: ID
  surveyId_gte: ID
  surveyId_lte: ID
  surveyId_in: [ID!]
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  survey: SurveyFilterType
}

type SurveyAnswerResultType {
  items: [SurveyAnswer!]!
  count: Int!
}`
)
