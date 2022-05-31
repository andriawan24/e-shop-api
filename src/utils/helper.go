package utils

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func APIResponse(message string, statusCode int, status string, data interface{}) Response {
	meta := Meta{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func FormatValidationErrors(err error) []string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, strings.Split(err.Error(), ":")[2])
	}

	return errors
}

func GetSecretKey() ([]byte, error) {
	err := godotenv.Load()
	secretKey := os.Getenv("SECRET_JWT_KEY")
	if err != nil {
		return []byte(secretKey), err
	}

	return []byte(secretKey), nil
}
