package api

import (
	"bytes"
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
			http.Error(w, "Error saving message. Try again later", http.StatusInternalServerError)
			return
		}

		err = broker.SendMessage(prod, &msg, id)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Error saving message. Try again later", http.StatusBadRequest)
			database.DeleteMessage(db, id)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}


func createStatsHandler(db *pgx.Conn) (http.HandlerFunc) {
	return func (w http.ResponseWriter, r *http.Request) {
		stats, err := database.GetStats(db)
		if err != nil {
			http.Error(w, "Error getting statistics", http.StatusInternalServerError)
			log.Print(err)
			return
		}
		var response bytes.Buffer
		json.NewEncoder(&response).Encode(stats)
		w.Write(response.Bytes())
	}
}