package config

import (
	"log"
	"os"
	"strconv"
)

type postgres struct {
	Host string
	Port int
	User string
	Pass string
}

func (p *postgres) fromEnv() *postgres {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("invalid db port")
	}
	p.Host = os.Getenv("DB_HOST")
	p.Port = port
	p.User = os.Getenv("DB_USER")
	p.Pass = os.Getenv("DB_PASS")
	return p
}
