package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", readAll)
	r.Get("/{docId}", read)
	r.Post("/create", createHandler)
	r.Put("/update", update)
	r.Delete("/delete/{docTitle}", delete)

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		panic(err.Error())
	}
}
