package models

// swagger:model
type InsertVeiculo struct {
	Placa string `json:"placa" validate:"required,len=7"`
}
