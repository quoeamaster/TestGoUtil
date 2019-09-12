package util

// definition of the Output plugin
type IOutputPlugin interface {
	Println(value...string)
	GetResponseMsg(msg string, code int) (responseMsg *ResponseMsg)
	GetRequestMsg(msg string, code int) (reqMsg *RequestMessage)
}


// RequestMessage related

type RequestMessage struct {
	Message string
	Code int
}

func NewRequestMessage() *RequestMessage {
	return new(RequestMessage)
}


// ResponseMsg

type ResponseMsg struct {
	Message string
	Code int
}