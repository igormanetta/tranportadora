package models

// swagger:model
type UpdateVeiculo struct {
	Placa string `json:"placa" validate:"required,len=7"`
}
