package controller

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-openapi/runtime/middleware"
)

type API struct {
	R *chi.Mux
}

func NewAPI() *API {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	return &API{R: r}
}

func (a *API) Listen() {

	a.R.Handle("/swagger/*", http.StripPrefix("/swagger", http.FileServer(http.Dir("./controller/swagger"))))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	a.R.Handle("/docs", sh)

	log.Println("listen on 8080")
	log.Println("doc: http://localhost:8080/docs")

	go func() {
		if err := http.ListenAndServe(":8080", a.R); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
