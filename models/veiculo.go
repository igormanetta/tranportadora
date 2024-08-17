package models

// swagger:model
type Veiculo struct {
	ID    string `json:"id"`
	Placa string `json:"placa" schema:"placa"`
}
