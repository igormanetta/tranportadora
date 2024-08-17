package models

import "github.com/google/uuid"

// swagger:model
type ReturnID struct {
	ID uuid.UUID `json:"id"`
}
