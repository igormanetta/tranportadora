package di

import (
	"log"
	"transportadora/controller"
)

func (d *DI) injectControllers() {
	d.Dig.Provide(controller.NewVeiculo)
	d.Dig.Provide(controller.NewMotorista)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		c.Routes()
	})

	if err != nil {
		log.Fatalf("error injection veiculo %v", err)
	}

	err = d.Dig.Invoke(func(c *controller.Motorista) {
		c.Routes()
	})

	if err != nil {
		log.Fatalf("error injection motorista %v", err)
	}
}
