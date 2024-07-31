package broker

import (
	"context"
	"log"
	"messagio/internal/database"

	"github.com/jackc/pgx/v5"
	"github.com/segmentio/kafka-go"
)

func RunConsumer(addr string, topic string, db *pgx.Conn) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: 	[]string{addr},
		Topic: 		topic,
		Partition: 	0,
		MinBytes: 	10e3,
		MaxBytes: 	10e6,
	})
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message from Kafka: ", err.Error())
		}
		
		id := string(m.Key)

		err = database.MarkMessageAsProcessed(db, id)
		if err != nil {
			log.Fatal("Failed to mark message as processed: ", err.Error())
		}
	}
}
