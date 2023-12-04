package rest

import (
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request
}

type ClientContentConfig struct {
	// ContentType specifies the wire format used to cmmunicate with the server.
	// This value will be set as the Accept header on requests made to the server if
	// AcceptContentTypes is not set, ad as the default content type on any object
	// sent to the server. If not set, "application/json" is used.
	ContentType string
}

type RESTClient struct {
	// base is teh root URL for all invocations of the client
	base *url.URL

	// content describes how a RESTClient encodes and decodes response.
	content ClientContentConfig

	// Set specific behavior of the client.
	Client *fasthttp.Client
}

func NewRestClient(baseURL *url.URL, config ClientContentConfig, client *fasthttp.Client) (*RESTClient, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}

	return &RESTClient{
		base:    &base,
		content: config,

		Client: client,
	}, nil
}
