package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var config model
var once sync.Once

type Config interface {
	Postgres() postgres
}

func Get() Config {
	once.Do(
		func() {
			if err := godotenv.Load("../.env"); err != nil {
				log.Fatalf("failed to load .env file: %s", err)
			}
			config.postgres = *new(postgres).fromEnv()
		},
	)
	return config
}

type model struct {
	postgres
}

func (m model) Postgres() postgres {
	return m.postgres
}
