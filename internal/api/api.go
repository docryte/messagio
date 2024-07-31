package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jackc/pgx/v5"
	"github.com/segmentio/kafka-go"
)

func Run(db *pgx.Conn, prod *kafka.Writer) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/message", createMessageHandler(db, prod))

	log.Fatal(http.ListenAndServe(":8080", r))
}
