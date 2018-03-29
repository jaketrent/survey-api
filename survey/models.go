package survey

type Survey struct {
	Id   string `json:"id"`
	Desc string `json:"desc"`
}

type Question struct {
	QuestionId string `json:"questionId"`
	SurveyId   string `json:"surveyId"`
	Prompt     string `json:"prompt"`
	Type       string `json:"type"`
}

type Answer struct {
	QuestionId string `json:"questionId"`
	Value      string `json:"value"`
}
