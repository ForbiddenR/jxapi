package api

import (
	"time"

	"github.com/valyala/fasthttp"
)

func Init(esamUrl, servicesUrl string, readTimeout, writeTimeout, maxIdleConnDuration time.Duration, maxConnsPerHost int) {
	EsamUrl, ServicesUrl = esamUrl, servicesUrl
	client = &fasthttp.Client{
		ReadTimeout:         readTimeout,
		WriteTimeout:        writeTimeout,
		MaxIdleConnDuration: maxIdleConnDuration,
		MaxConnsPerHost:     maxConnsPerHost,
	}
}
