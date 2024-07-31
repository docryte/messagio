package api

import (
	"encoding/json"
	"log"
	"net/http"

	"messagio/internal/broker"
	"messagio/internal/database"
	"messagio/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/segmentio/kafka-go"
)

func createMessageHandler(db *pgx.Conn, prod *kafka.Writer) (http.HandlerFunc) {
	return func (w http.ResponseWriter, r *http.Request) {
		var msg models.Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Bad message", 400)
			return
		}
		id, err := database.SaveMessage(db, &msg)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Error saving message. Try again later", 500)
			return
		}
		err = broker.SendMessage(prod, &msg, id)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Error saving message. Try again later", 500)
			database.DeleteMessage(db, id)
			return
		}
	}
}
