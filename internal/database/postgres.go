package database

import (
	"github.com/jackc/pgx/v5"

	"context"
	"log"

	"messagio/internal/config"
	"messagio/internal/models"
)

var db *pgx.Conn

func initPostgres(cfg *config.Config) {
	var err error
	db, err = pgx.Connect(context.Background(), cfg.PostgresUrl)
	if err != nil {
		log.Fatal("Error connecting postgres database: ", err.Error())
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Error resolving postgres database: ", err.Error())
	}
}

func SaveMessage(msg *models.Message) (id int, err error) {
	err = db.QueryRow(
		context.Background(),
		"INSERT INTO messages (content) VALUES ($1) RETURNING id",
		msg.Content,
		false,
	).Scan(&id)
	return
}

func MarkMessageAsProcessed(id int) (err error) {
	_, err = db.Exec(context.Background(), "UPDATE messages SET processed_at = NOW() WHERE id = $1", id)
	return
}
