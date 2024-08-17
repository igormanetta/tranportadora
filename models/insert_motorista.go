package models

// swagger:model
type InsertMotorista struct {
	Nome string `json:"nome" validate:"required,max=255,min=3"`
}
