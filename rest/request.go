package rest

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ForbiddenR/jxapi/apierrors"
	"github.com/valyala/fasthttp"
)

type Request struct {
	c *RESTClient

	request *fasthttp.Request
	path    string

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

func (r *Request) RequestURI(uri string) *Request {
	if r.err != nil {
		return r
	}
	r.path = uri
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

func (r *Request) Do(ctx context.Context) Result {
	finalURL := &url.URL{}
	if r.c.base != nil {
		*finalURL = *r.c.base
	}
	finalURL.Path = r.path
	r.request.SetRequestURI(finalURL.String())
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(r.request)
	}()
	err := r.c.Client.DoTimeout(r.request, resp, 3*time.Second)
	if err != nil {
		return Result{err: apierrors.GetFailedRequestDoTimeoutError(err)}
	}
	if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
		if statusCode == fasthttp.StatusNotFound {
			return Result{err: ErrNotFound}
		}
		return Result{err: ErrServicesException}
	}
	respBody := resp.Body()
	if len(respBody) == 0 {
		return Result{err: ErrBodyIsNil}
	}
	body := make([]byte, len(respBody))
	copy(body, respBody)
	return Result{body: body}
}

// Error returns any error encountered construcing the reuest, if any.
func (r *Request) Error() error {
	return r.err
}

type Result struct {
	body []byte
	err  error
}

func (r Result) Into(obj interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal(r.body, obj)
}
