package main

import (
	"messagio/internal/api"
	"messagio/internal/broker"
	"messagio/internal/config"
	"messagio/internal/database"
)

func main() {
	cfg := config.Read()
	db := database.GetConnection(cfg.PostgresUrl)
	prod := broker.GetProducer(cfg.KafkaUrl)
	go broker.RunConsumer(cfg.KafkaUrl, db)
	api.Run(db, prod)
}