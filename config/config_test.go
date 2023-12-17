package config_test

import (
	"os"
	"testing"
	"web-service-gin/config"
)

func TestGet_Postgres(t *testing.T) {
	// Happy test case
	cfg := config.Get()
	if cfg == nil {
		t.Fatal("expected config, got nil")
	}
	if cfg.Postgres().Host != "localhost" {
		t.Errorf("expected postgres.host = %s, got %s", "localhost", cfg.Postgres().Host)
	}
	// Test case 1
	err := os.Setenv("DB_HOST", "updated-db-host-value-using-os.Setenv") // Developer mistake in a goroutine!
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Get config again in another goroutine
	cfg3 := config.Get()
	if cfg3.Postgres().Host != "localhost" {
		// if comment once.Do(...) lines in config.Get() method in config.go
		t.Errorf("expected postgres.host = %s, got %s", "localhost", cfg3.Postgres().Host)
	}
}
