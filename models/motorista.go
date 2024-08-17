package models

// swagger:model
type Motorista struct {
	ID      string   `json:"id"`
	Nome    string   `json:"nome" schema:"nome"`
	Veiculo *Veiculo `json:"veiculo"`
}
