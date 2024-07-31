package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	PostgresUrl	string	`env:"POSTGRES_URL"`
	KafkaUrl	string	`env:"KAFKA_URL"`
}

func Read() *Config {
	cfg := &Config{}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
