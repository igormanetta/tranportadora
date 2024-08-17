package models

// swagger:model
type ListVeiculo struct {
	Pagination PaginationResponse
	Data       *[]Veiculo
}
