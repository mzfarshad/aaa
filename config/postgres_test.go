package config

import (
	"os"
	"strconv"
	"testing"
)

func TestPostgres_fromEnv(t *testing.T) {
	// TODO: @Farshad
	// 1. Use os.Setenv to set "DB_HOST","DB_PORT","DB_USER", and "DB_PASS" to some values
	// 2. Assign a variable (for example, p) to new(postgres).fromEnv()
	// 3. Test failed if p.Host != os.Getenv("DB_HOST"), and so on...
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "me")
	os.Setenv("DB_PASS", "pass123")
	os.Setenv("DB_NAME", "aaa")
	os.Setenv("DB_TIMEZONE", "Asia/Tehran")

	p, err := new(postgres).fromEnv()
	if err != nil {
		t.Error(err)
	}
	if p.Host != os.Getenv("DB_HOST") {
		t.Errorf("expected p.host = localhost, got %s ", os.Getenv("DB_HOST"))
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		t.Fatal(err.Error())
	}
	if p.Port != port {
		t.Errorf("expected p.port = 5432, got %d ", port)
	}
	if p.User != os.Getenv("DB_USER") {
		t.Errorf("expected p.user = me, got %s ", os.Getenv("DB_USER"))
	}
	if p.Pass != os.Getenv("DB_PASS") {
		t.Errorf("expected p.pass = pass123, got %s ", os.Getenv("DB_PASS"))
	}
	if p.Name != os.Getenv("DB_NAME") {
		t.Errorf("expected p.name = aaa, got %s ", os.Getenv("DB_NAME"))
	}
	if p.TimeZone != os.Getenv("DB_TIMEZONE") {
		t.Errorf("expected p.timezone = Asia/Tehran, got %s ", os.Getenv("DB_TIMEZONE"))
	}
}
