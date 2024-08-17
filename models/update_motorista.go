package models

// swagger:model
type UpdateMotorista struct {
	Nome string `json:"nome" validate:"required,max=255"`
}
