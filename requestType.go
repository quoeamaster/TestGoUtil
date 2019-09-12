package util

type RequestMessage struct {
	Message string
	Code int
}

func NewRequestMessage() *RequestMessage {
	return new(RequestMessage)
}
