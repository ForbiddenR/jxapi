package apierrors

import "fmt"

const (
	FailedRequest           = "SendRequest's Error: Send request %s to url %s failed: error: %v"
	FailedRequestMarshal    = "SendRequest's json.Marsh's Error: marshal request %v to json failed. error: %v"
	FailedRequestSend       = "SendRequest's Do Timeout's Error: do with time failed: error: %v"
	FailedResponseUnmarshal = "json.Unmarshal's Error: Send request: %s to url: %s, but unmarshal response: %s failed. error: %v"
)

func GetFailedResponseUnmarshalError(url string, req, resp []byte, err error) error {
	return fmt.Errorf(FailedResponseUnmarshal, req, url, resp, err)
}

func GetFailedRequestMarshalError(req []byte, err error) error {
	return fmt.Errorf(FailedRequestMarshal, req, err)
}

func GetFailedRequestDoTimeoutError(err error) error {
	return fmt.Errorf(FailedRequestSend, err)
}

func GetFailedRequestError(url string, req []byte, err error) error {
	return fmt.Errorf(FailedRequest, req, url, err)
}
