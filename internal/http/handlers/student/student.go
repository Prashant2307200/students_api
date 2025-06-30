package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Prashant2307200/students-api/internal"
	"github.com/Prashant2307200/students-api/internal/types"
	"github.com/Prashant2307200/students-api/internal/utils/response"
	"github.com/go-playground/validator"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("request body is empty")))
			return
		}

		if err := validator.New().Struct(&student); err != nil {
			slog.Error("Validation error", slog.Any("error", err))
			if ve, ok := err.(validator.ValidationErrors); ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(ve))
			} else {
				response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			}
			return
		}

		slog.Info("Creating a student", slog.Any("student", student))
		id, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			slog.Error("Failed to create student", slog.Any("error", err))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		slog.Info("Student created successfully", slog.Int64("id", id))
		w.Header().Set("Location", r.URL.Path+"/"+fmt.Sprint(id))

		response.WriteJson(w, http.StatusOK, map[string]int64{"id": id})
	}
}
