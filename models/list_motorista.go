package models

// swagger:model
type ListMotorista struct {
	Pagination PaginationResponse
	Data       *[]Motorista
}
