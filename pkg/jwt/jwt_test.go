package jwt_test

import (
	"log"
	"os"
	"testing"
	"web-service-gin/models"
	"web-service-gin/pkg/jwt"
)

func TestNewAccessToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET_KEY", "256-bit-secret")
	if err != nil {
		t.Fatalf(err.Error())
	}
	token, err := jwt.NewAccessToken("x@y.z", models.UserTypeAdmin)
	if err != nil {
		t.Fatalf("expected new access token, got err: %s", err.Error())
	}
	if token == "" {
		t.Errorf("expected access token, got empty token")
	}
	log.Println(token)
}
