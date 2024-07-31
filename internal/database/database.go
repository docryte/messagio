package database

import "messagio/internal/config"

func Init(cfg *config.Config) {
	initPostgres(cfg)
}