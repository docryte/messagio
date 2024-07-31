package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/message", createMessage)

	log.Fatal(http.ListenAndServe(":8080", r))
}
