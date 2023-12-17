package config

import (
	"log"
	"sync"
)

var config model
var once sync.Once

type Config interface {
	Postgres() *postgres
	JWT() *jwt
}

func Get() Config {
	once.Do(
		func() {
			psql, err := new(postgres).fromEnv()
			if err != nil {
				log.Println(err)
			}
			config.postgres = *psql
			config.jwt = *new(jwt).fromEnv()
		},
	)
	return config
}

type model struct {
	postgres
	jwt
}

func (m model) Postgres() *postgres {
	return &m.postgres
}

func (m model) JWT() *jwt {
	return &m.jwt
}
