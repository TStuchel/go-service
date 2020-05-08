package common

type ErrorDTO struct {
	Url string `json:"url"`
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
	Type string `json:"type"`
}