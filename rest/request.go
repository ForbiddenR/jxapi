package rest

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ForbiddenR/jxapi/apierrors"
	"github.com/valyala/fasthttp"
)

type Request struct {
	c *RESTClient

	request *fasthttp.Request

	// output
	err error
}

func NewRequest(c *RESTClient) *Request {
	r := &Request{
		c:       c,
		request: fasthttp.AcquireRequest(),
	}
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.request.Header.SetMethod(verb)
	return r
}

func (r *Request) RequestURI(url string) *Request {
	if r.err != nil {
		return r
	}
	r.request.SetRequestURI(url)
	return r
}

func (r *Request) SetHeader(header map[string]string) *Request {
	for k, v := range header {
		r.request.Header.Set(k, v)
	}
	return r
}

func (r *Request) Body(body interface{}) *Request {
	if r.err != nil {
		return r
	}
	raw, err := json.Marshal(body)
	if err != nil {
		r.err = err
		return r
	}
	r.request.SetBodyRaw(raw)
	return r
}

func (r *Request) Do(ctx context.Context) error {
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(r.request)
	}()
	err := r.c.Client.DoTimeout(r.request, resp, 3*time.Second)
	if err != nil {
		return apierrors.GetFailedRequestDoTimeoutError(err)
	}
	if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
		if statusCode == fasthttp.StatusNotFound {
			return ErrNotFound
		}
		return ErrServicesException
	}
	respBody := resp.Body()
	if len(respBody) == 0 {
		return ErrBodyIsNil
	}
	return nil
}
