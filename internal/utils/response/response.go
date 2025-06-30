package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errors []string
	for _, err := range errs {
		switch err.ActualTag() {

		case "required":
			errors = append(errors, fmt.Sprintf("Field %s is required", err.Field()))
		default:
			errors = append(errors, fmt.Sprintf("Field %s is %s", err.Field(), err.ActualTag()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  "Validation failed: " + strings.Join(errors, ", "),
	}
}
