package presenter

type Response struct {
	Data      any      `json:"data"`
	IsSuccess bool     `json:"is_success"`
	Messages  []string `json:"messages,omitempty"`
}

func (r Response) AppendMessage(msg string) Response {
	if msg == "" { // Do not append empty string
		return r
	}
	// Append new message
	messages := r.Messages
	messages = append(messages, msg)
	return Response{
		Data:      r.Data,
		IsSuccess: r.IsSuccess,
		Messages:  messages,
	}
}

func NewSuccess(data any) Response {
	return Response{
		Data:      data,
		IsSuccess: true,
	}
}

func NewFailed(message string) Response {
	return Response{
		IsSuccess: false,
		Messages:  []string{message},
	}
}
