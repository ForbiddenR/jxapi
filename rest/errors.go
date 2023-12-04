package rest

import "errors"

var (
	ErrBodyIsNil = errors.New("body is nil")
	// service端发生异常导致未返回数据
	ErrServicesException = errors.New("services exception")
	// 404 not found error
	ErrNotFound = errors.New("not found")
)
