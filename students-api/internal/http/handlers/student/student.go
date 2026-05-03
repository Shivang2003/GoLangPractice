package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id") //this is a method provided by the http package to extract path parameters from the request URL. It takes the name of the parameter as an argument and returns its value as a string.
		slog.Info("Getting a student by id", slog.String("path", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			slog.Error(err.Error())
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id")))
			return
		}

		student, err := storage.GetStudentByID(intId)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)

	}
}
