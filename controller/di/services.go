package di

import (
	"transportadora/services"

	"go.uber.org/dig"
)

func (d *DI) injectServices() {
	err := d.Dig.Provide(services.NewVeiculo, dig.As(new(services.Veiculo)))
	if err != nil {
		panic(err)
	}
	d.Dig.Provide(services.NewMotorista, dig.As(new(services.Motorista)))
}
