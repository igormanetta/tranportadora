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

type Veiculo struct {
	API     *API
	service services.Veiculo
}

func NewVeiculo(api *API, service services.Veiculo) *Veiculo {
	return &Veiculo{
		API:     api,
		service: service,
	}
}

func (c *Veiculo) HandleInsertVeiculo(w http.ResponseWriter, r *http.Request) {
	var dto models.InsertVeiculo
	if err := UnmarshalAndValidate(r.Body, &dto); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	id, err := c.service.InsertVeiculo(r.Context(), dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	returnValue := models.ReturnID{
		ID: id,
	}

	CreatedResponse(w, r, returnValue)
}

func (c *Veiculo) HandleUpdateVeiculo(w http.ResponseWriter, r *http.Request) {
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

	var dto models.UpdateVeiculo
	if err := UnmarshalAndValidate(r.Body, &dto); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	id, err = c.service.UpdateVeiculo(r.Context(), id, dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	returnValue := models.ReturnID{
		ID: id,
	}

	OkResponse(w, r, returnValue)
}

func (c *Veiculo) HandleDeleteVeiculo(w http.ResponseWriter, r *http.Request) {
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

	if err := c.service.DeleteVeiculo(r.Context(), id); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	NoContentResponse(w, r)
}

func (c *Veiculo) HandleGetVeiculo(w http.ResponseWriter, r *http.Request) {
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

	veiculo, err := c.service.GetVeiculo(r.Context(), id)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	OkResponse(w, r, veiculo)
}

func (c *Veiculo) HandleListVeiculos(w http.ResponseWriter, r *http.Request) {
	var dto models.SearchVeiculo

	err := schema.NewDecoder().Decode(&dto, r.URL.Query())
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	veiculos, err := c.service.ListVeiculos(r.Context(), dto)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	OkResponse(w, r, veiculos)
}

func (c *Veiculo) Routes() {
	c.API.R.Post("/veiculo", c.HandleInsertVeiculo)
	c.API.R.Put("/veiculo/{id}", c.HandleUpdateVeiculo)
	c.API.R.Delete("/veiculo/{id}", c.HandleDeleteVeiculo)

	c.API.R.Get("/veiculo/{id}", c.HandleGetVeiculo)
	c.API.R.Get("/veiculo", c.HandleListVeiculos)
}
