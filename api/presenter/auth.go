package presenter

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Access string `json:"access_token"`
}
