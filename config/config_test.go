package config_test

import (
	"os"
	"testing"
	"web-service-gin/config"
)

func TestGet(t *testing.T) {
	cfg := config.Get()
	if cfg.Postgres.Host != "localhost" {
		t.Errorf("expected postgres.host = %s, got %s", "localhost", cfg.Postgres.Host)
	}

	cfg.Postgres.Host = "something-else" // Developer mistake in a goroutine!
	// Get config again in another goroutine
	cfg2 := config.Get()
	if cfg2.Postgres.Host != "localhost" {
		t.Errorf("expected postgres.host = %s, got %s", "localhost", cfg2.Postgres.Host)
	}

	err := os.Setenv("DB_HOST", "updated-db-host-env") // Developer mistake in a goroutine!
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Get config again in another goroutine
	cfg3 := config.Get()
	if cfg3.Postgres.Host != "localhost" {
		// if comment once.Do(...) lines in config.Get() method in config.go
		t.Errorf("expected postgres.host = %s, got %s", "localhost", cfg3.Postgres.Host)
	}
}
