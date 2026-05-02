package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shivang2003/students-api/internal/storage"
	"github.com/shivang2003/students-api/internal/types"
	"github.com/shivang2003/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) { //IF BODY IS EMPTY
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation

		validate := validator.New()

		if err := validate.Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)

			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validateErrs))
			return
		}

		lastInsertID, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		slog.Info("user created with id, ", slog.String("ID:", fmt.Sprint(lastInsertID)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastInsertID})
	}
}
