package rest

import "errors"

var (
	ErrBodyIsNil = errors.New("body is nil")
	// service-caused error
	ErrServicesException = errors.New("services exception")
	// 404 not found error
	ErrNotFound = errors.New("not found")
)
