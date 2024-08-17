package main

import (
	"log"
	"transportadora/controller"
	"transportadora/controller/di"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using default values")
	}

	d := di.New()
	d.Inject(false)

	defer d.Close(false)

	err := d.Dig.Invoke(func(s *controller.API) {
		s.Listen()
	})

	if err != nil {
		log.Fatalf("erro ao subir server: %v", err)
	}
}
