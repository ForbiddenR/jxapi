package apierrors_test

import (
	"errors"
	"github.com/ForbiddenR/jxapi/apierrors"
	"testing"
)

func TestGetFailedRequestMarshalError(t *testing.T) {
	result := apierrors.GetFailedRequestMarshalError([]byte("requestId: 12"), errors.New("failed to marshal"))
	t.Log(result.Error())
}

func TestGetFailedRequestDoTimeoutError(t *testing.T) {
	result := apierrors.GetFailedRequestDoTimeoutError(errors.New("failed to connect to server"))
	t.Log(result.Error())
}

func TestGetFailedResponseUnmarshalError(t *testing.T) {
	result := apierrors.GetFailedResponseUnmarshalError("https://google.com", []byte("requestId: 12"), []byte("responseId: 13"),
		errors.New("ambiguous response format"))
	t.Log(result.Error())
}
