package models

// swagger:model
type PaginationResponse struct {
	RecordPerPage int `json:"record_per_page"`
	CurrentPage   int `json:"current_page"`
	TotalRecord   int `json:"total_record"`
	TotalPage     int `json:"total_page"`
}

// swagger:model
type PaginationRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
