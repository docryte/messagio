package main

import (
	"messagio/internal/api"
	"messagio/internal/config"
	"messagio/internal/database"
)

func main() {
	cfg := config.Read()
	database.Init(cfg)
	api.Init()
}