package models

import "time"

type Statistics struct {
	TotalProccessedMessages  	int64
	LastProcessedMessageTime 	*time.Time
	AverageProcessingTimeMs  	*float64
	QueuedMessages				int64
}