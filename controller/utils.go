package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"transportadora/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

func UnmarshalAndValidate(body io.ReadCloser, dest interface{}) error {
	err := json.NewDecoder(body).Decode(dest)
	if err != nil {
		return err
	}

	validate := validator.New()
	return validate.Struct(dest)
}

func OkResponse(w http.ResponseWriter, r *http.Request, obj interface{}) {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func NoContentResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, obj interface{}) {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	responseMessage(w, err, http.StatusBadRequest)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, pgx.ErrNoRows) {
		responseMessage(w, err, http.StatusNotFound)
		return
	}

	if errors.Is(err, utils.ErrMotAssociado) {
		responseMessage(w, err, http.StatusConflict)
		return
	}
}

func responseMessage(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
}
