package config

import "os"

type jwt struct {
	SecretKey string
}

func (j *jwt) fromEnv() *jwt {
	j.SecretKey = os.Getenv("JWT_SECRET_KEY")
	return j
}
