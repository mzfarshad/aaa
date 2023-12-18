package presenter

import (
	"web-service-gin/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) From(user *models.User) *LoginRequest {
	if user == nil {
		return nil
	}
	if l == nil {
		return new(LoginRequest).From(user)
	}
	l.Email = user.Email
	l.Password = user.Password
	return l
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Access string `json:"access_token"`
}
