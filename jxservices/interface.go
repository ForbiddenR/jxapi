package jxservices

// Request interface needs to be implemented by all api.
type Request interface {
	GetName() string
}

// Response will be implemented by all response struct.
type Response interface {
	GetStatus() int
	GetMsg() string
}

// CallbackRequest is needed to be implemented by all the callback api.
type CallbackRequest interface {
	Request
}

type ReusedConfig struct {
	Sn       string
	Protocol *Protocol
	Pod      string
	MsgID    string
}