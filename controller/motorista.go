package controller

import (
	"errors"
	"net/http"
	"transportadora/models"
	"transportadora/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

type Motorista struct {
	API     *API
	service services.Motorista
}

func NewMotorista(api *API, service services.Motorista) *Motorista {
	return &Motorista{
		API:     api,
		service: service,
	}
}

func (c *Motorista) HandleInsertMotorista(w http.ResponseWriter, r *http.Request) {
	var dto models.InsertMotorista
	if err := UnmarshalAndValidate(r.Body, &dto); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	id, err := c.service.InsertMotorista(r.Context(), dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	returnValue := models.ReturnID{
		ID: id,
	}

	CreatedResponse(w, r, returnValue)
}

func (c *Motorista) HandleUpdateMotorista(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	if rawID == "" {
		BadRequestResponse(w, r, errors.New("missing id"))
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	var dto models.UpdateMotorista
	if err := UnmarshalAndValidate(r.Body, &dto); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	id, err = c.service.UpdateMotorista(r.Context(), id, dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	returnValue := models.ReturnID{
		ID: id,
	}

	OkResponse(w, r, returnValue)
}

func (c *Motorista) HandleDeleteMotorista(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	if rawID == "" {
		BadRequestResponse(w, r, errors.New("missing id"))
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	if err := c.service.DeleteMotorista(r.Context(), id); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	NoContentResponse(w, r)
}

func (c *Motorista) HandleGetMotorista(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	if rawID == "" {
		BadRequestResponse(w, r, errors.New("missing id"))
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	motorista, err := c.service.GetMotorista(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	OkResponse(w, r, motorista)
}

func (c *Motorista) HandleListMotoristas(w http.ResponseWriter, r *http.Request) {
	var dto models.SearchMotorista

	err := schema.NewDecoder().Decode(&dto, r.URL.Query())
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	listMotorista, err := c.service.ListMotoristas(r.Context(), dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	OkResponse(w, r, listMotorista)
}

func (c *Motorista) HandleSetMotoristaVeiculo(w http.ResponseWriter, r *http.Request) {
	rawMotoristaID := chi.URLParam(r, "motoristaID")
	if rawMotoristaID == "" {
		BadRequestResponse(w, r, errors.New("missing motoristaID"))
		return
	}

	motoristaID, err := uuid.Parse(rawMotoristaID)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	rawVeiculoID := chi.URLParam(r, "veiculoID")
	if rawVeiculoID == "" {
		BadRequestResponse(w, r, errors.New("missing veiculoID"))
		return
	}

	veiculoID, err := uuid.Parse(rawVeiculoID)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	id, err := c.service.SetMotoristaVeiculo(r.Context(), motoristaID, veiculoID)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	returnValue := models.ReturnID{
		ID: id,
	}

	OkResponse(w, r, returnValue)
}

func (c *Motorista) Routes() {
	c.API.R.Post("/motorista", c.HandleInsertMotorista)
	c.API.R.Put("/motorista/{id}", c.HandleUpdateMotorista)
	c.API.R.Delete("/motorista/{id}", c.HandleDeleteMotorista)

	c.API.R.Get("/motorista/{id}", c.HandleGetMotorista)
	c.API.R.Get("/motorista", c.HandleListMotoristas)

	c.API.R.Patch("/motorista/{motoristaID}/veiculo/{veiculoID}", c.HandleSetMotoristaVeiculo)
}
