package api

import (
	"encoding/json"
	"log"
	"net/http"

	"messagio/internal/database"
	"messagio/internal/models"
)

func createMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Bad message", 400)
		return
	}
	_, err = database.SaveMessage(&msg)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error saving message. Try again later", 500)
		return
	}
}
