package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Postgres struct {
	Host string
	Port int
	User string
	Pass string
}
type Config struct {
	Postgres
}

func ConnectDataBase() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %s", err)
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("invalid port")
	}
	config := Config{
		Postgres{
			Host: os.Getenv("HOST"),
			Port: port,
			User: os.Getenv("USER_NAME"),
			Pass: os.Getenv("PASS"),
		},
	}
	return config
}
