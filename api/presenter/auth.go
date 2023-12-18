package presenter

import (
	"errors"
	"regexp"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Validate validates the given reciever's email and password.
// This is just for learning purposes: Now, we used gin's binding tag
// in request structs to validate.
func (req SignInRequest) Validate() error {
	if req.Email == "" {
		return errors.New("email is required")
	}
	// TODO: @Farshad
	//Learn more about regular expressions in https://regexone.com
	mathed, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, req.Email)
	if err != nil || !mathed {
		return errors.New("invalid email")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Access string `json:"access_token"`
}
