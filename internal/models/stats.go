package models

import "time"

type Statistics struct {
	TotalProccessedMessages  	int
	LastProcessedMessageTime 	time.Time
	AverageProcessingTimeMs  	float64
	QueuedMessages				int
}