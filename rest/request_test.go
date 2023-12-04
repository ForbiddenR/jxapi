package rest

import (
	"net/url"
	"testing"
)

func Test(t *testing.T) {
	finalURL := &url.URL{}
	finalURL.Host= "http://localhost"
	finalURL.Path="test"
}