package models

type GeneralResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error"`
}

type CompletionResponse struct {
	Status     int        `json:"status"`
	Message    string     `json:"message"`
	Completion Completion `json:"completion"`
}

type Completion struct {
	Id       string `json:"completionId"`
	UserId   string `json:"userId"`
	Question string `json:"question"`
	Response string `json:"response"`
}
