package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once
var Config config

type config struct {
	Postgres postgres
}

func Get() config {
	once.Do(
		func() {
			if err := godotenv.Load("../.env"); err != nil {
				log.Fatalf("failed to load .env file: %s", err)
			}
			Config.Postgres = *new(postgres).fromEnv()
		},
	)
	return Config
}
