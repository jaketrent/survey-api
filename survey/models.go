package survey

type Survey struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type Question struct {
	Id           int    `json:"id"`
	SurveyId     int    `json:"surveyId"`
	Prompt       string `json:"prompt"`
	QuestionType string `json:"questionType"`
}

type Answer struct {
	Id         int    `json:"id"`
	QuestionId string `json:"questionId"`
	Value      string `json:"value"`
	UserInfo   string `json:"userInfo"`
}
