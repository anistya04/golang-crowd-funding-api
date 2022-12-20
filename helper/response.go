package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

func ApiResponse(message string, code int, status string, data any) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func ErrorResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
