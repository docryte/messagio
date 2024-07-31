package broker

import (
	"context"
	"messagio/internal/models"

	"github.com/segmentio/kafka-go"
)

func GetProducer(addr string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:		kafka.TCP(addr),
		Topic:		topic,
		Balancer:	&kafka.LeastBytes{},
	}
}

func SendMessage(prod *kafka.Writer, msg *models.Message, id string) (err error) {
	err = prod.WriteMessages(context.Background(), kafka.Message{
		Key: 	[]byte(id),
		Value: 	[]byte(msg.Content),
	})
	return
}