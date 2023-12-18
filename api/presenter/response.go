package presenter

type Response struct {
	Data      any      `json:"data"`
	IsSuccess bool     `json:"is_success"`
	Messages  []string `json:"messages,omitempty"`
}

func (r Response) AppendMessages(messages ...string) Response {
	if len(messages) == 0 { // Do not append empty slice
		return r
	}
	// Append new message
	responseMessages := r.Messages
	responseMessages = append(responseMessages, messages...)
	return Response{
		Data:      r.Data,
		IsSuccess: r.IsSuccess,
		Messages:  responseMessages,
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
