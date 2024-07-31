package database

import (
	"github.com/jackc/pgx/v5"

	"context"
	"log"

	"messagio/internal/models"
)

func GetConnection(addr string) (db *pgx.Conn) {
	var err error
	db, err = pgx.Connect(context.Background(), addr)
	if err != nil {
		log.Fatal("Error connecting postgres database: ", err.Error())
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Error resolving postgres database: ", err.Error())
	}
	return
}

func SaveMessage(db *pgx.Conn, msg *models.Message) (id string, err error) {
	err = db.QueryRow(
		context.Background(),
		"INSERT INTO messages (content) VALUES ($1) RETURNING id",
		msg.Content,
		false,
	).Scan(&id)
	return
}

func DeleteMessage(db *pgx.Conn, id string) (err error) {
	_, err = db.Exec(
		context.Background(),
		"DELETE FROM messages WHERE id = $1",
		id,
	)
	return
}

func MarkMessageAsProcessed(db *pgx.Conn, id string) (err error) {
	_, err = db.Exec(context.Background(), "UPDATE messages SET processed_at = NOW() WHERE id = $1", id)
	return
}

func GetStats(db *pgx.Conn) (stats models.Statistics, err error) {
	err = db.QueryRow(
		context.Background(),
		`
			SELECT
				COUNT (*) AS total_processed_messages,
				MAX(processed_at) AS last_processed_message_time,
				AVG(EXTRACT(EPOCH FROM (processed_at - created_at)) * 1000) AS average_processing_time_ms,
				(SELECT COUNT(*) FROM messages WHERE processed_at = null) AS queued_messages
			FROM
				messages
			WHERE
				processed_at != null
		`,
	).Scan(
		&stats.TotalProccessedMessages,
		&stats.LastProcessedMessageTime,
		&stats.AverageProcessingTimeMs,
		&stats.QueuedMessages,
	)
	return
} 