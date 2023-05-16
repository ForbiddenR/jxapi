package apierrors

import (
	"fmt"
)

type CallbackErrorCodeType string

const (
	// CallbackErrorCodeTypeOffline happens when the station is offline.
	CallbackErrorCodeTypeOffline CallbackErrorCodeType = "Offline"
	// 请求参数错误
	CallbackErrorCodeTypeRequestParamsIncorrect CallbackErrorCodeType = "RequestParamsIncorrect"
	// CallbackErrorCodeTypeTimeout happens when station fails to reply the request within specified time.
	CallbackErrorCodeTypeTimeout CallbackErrorCodeType = "Timeout"
	// 桩回复未实现该命令
	CallbackErrorCodeTypeNotImplemented CallbackErrorCodeType = "NotImplemented"
	// CallbackErrorCodeTypeInternalError 桩回复内部错误
	CallbackErrorCodeTypeInternalError CallbackErrorCodeType = "InternalError"
	// Requested Action is recognized but not supported by the cs.
	CallbackErrorCodeTypeNotSupported CallbackErrorCodeType = "NotSupported"
	// CallbackErrorCodeTypeSecurityError happens when during the processing of Action a security issue occurred
	//preventing receiver from completing the Action successfully.
	CallbackErrorCodeTypeSecurityError CallbackErrorCodeType = "SecurityError"
	// CallbackErrorCodeTypePayloadError happens when some error is caused by wrong payload.
	CallbackErrorCodeTypePayloadError CallbackErrorCodeType = "PayloadError"
	// Any other error not covered by the previous ones.
	CallbackErrorCodeTypeGenericError CallbackErrorCodeType = "GenericError"
)

type CallbackError struct {
	code  CallbackErrorCodeType
	inner error
}

func (c *CallbackError) Code() CallbackErrorCodeType {
	return c.code
}

func (c *CallbackError) Error() string {
	return c.inner.Error()
}

func NewCallbackErrorOffline(id, command string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeOffline,
		inner: fmt.Errorf("command:%s cannot send to %s because of offline", command, id),
	}
}

func NewCallbackErrorRequestParamsIncorrect(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeOffline,
		inner: fmt.Errorf("command:%s cannot send to %s because of params incorrect: %s", command, id, reason),
	}
}

func NewCallbackErrorTimeout(id, command string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeTimeout,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of timeout", command, id),
	}
}

func NewCallbackErrorNotImplemented(id, command string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeNotImplemented,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of not implemented", command, id),
	}
}

func NewCallbackErrorInternalError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeInternalError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of InternalError: %s", command, id, reason),
	}
}

func NewCallbackErrorSupported(id, command string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeNotSupported,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of being Supported", command, id),
	}
}

func NewCallbackErrorGenericError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeGenericError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of a generic error: %s", command, id, reason),
	}
}

func NewCallbackErrorPayloadError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypePayloadError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of wrong payload: %s", command, id, reason),
	}
}

func NewCallbackErrorResponsePayloadError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypePayloadError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of empty payload in DataTransfer: %s", command, id, reason),
	}
}

func NewCallbackErrorWrongParsedPayloadError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypePayloadError,
		inner: fmt.Errorf("command:%s send to %s but the parsing process has failed in: %s", command, id, reason),
	}
}

//func NewCallbackErrorSecurityError(id, command, reason string) error {

func NewCallbackErrorSecurityError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypeSecurityError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of a security issue；%s", command, id, reason),
	}
}

func NewCallbackErrorInvalidConfigurationError(id, command, reason string) *CallbackError {
	return &CallbackError{
		code:  CallbackErrorCodeTypePayloadError,
		inner: fmt.Errorf("command:%s send to %s but callback failed because of a invalid configuration: %s", command, id, reason),
	}
}
