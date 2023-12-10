package presenter

type Response struct {
	Data      any      `json:"data"`
	IsSuccess bool     `json:"is_success"`
	Messages  []string `json:"messages,omitempty"`
}
