package di

import (
	pgstore "transportadora/infra/pgstore/repositories"
	"transportadora/infra/repositories"

	"go.uber.org/dig"
)

func (d *DI) injectRepos() {
	d.Dig.Provide(pgstore.NewVeiculo, dig.As(new(repositories.Veiculo)))
	d.Dig.Provide(pgstore.NewMotorista, dig.As(new(repositories.Motorista)))
}
