package models

type ErrorHandler struct {
	StatusCode int    `json:"statusCode"`
	Mensaje    string `json:"mensaje"`
}
