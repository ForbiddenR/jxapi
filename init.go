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

type options struct {
	readTimeout         time.Duration
	writeTimeout        time.Duration
	maxIdleConnDuration time.Duration
	maxConnsPerHost     int
}

type Option func(options *options) error

func WithReadTimeout(readtimeout time.Duration) Option {
	return func(options *options) error {
		options.readTimeout = readtimeout
		return nil
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(options *options) error {
		options.writeTimeout = writeTimeout
		return nil
	}
}

func WithMaxIdleConnDuration(maxIdleConnDuration time.Duration) Option {
	return func(options *options) error {
		options.maxIdleConnDuration = maxIdleConnDuration
		return nil
	}
}

func WithMaxConnsPerHost(maxConnsPerHost int) Option {
	return func(options *options) error {
		options.maxConnsPerHost = maxConnsPerHost
		return nil
	}
}

func InitApi(esamUrl, servicesUrl string, opts ...Option) (err error) {
	options := options{}
	for _, opt := range opts {
		err = opt(&options)
		if err != nil {
			return err
		}
	}

	if options.readTimeout == 0 {
		options.readTimeout = 10 * time.Second
	}
	if options.writeTimeout == 0 {
		options.readTimeout = 10 * time.Second
	}
	if options.maxIdleConnDuration == 0 {
		options.maxIdleConnDuration = 10 * time.Second
	}
	if options.maxConnsPerHost == 0 {
		options.maxConnsPerHost = 5000
	}

	EsamUrl, ServicesUrl = esamUrl, servicesUrl
	client = &fasthttp.Client{
		ReadTimeout:         options.readTimeout,
		WriteTimeout:        options.writeTimeout,
		MaxIdleConnDuration: options.maxIdleConnDuration,
		MaxConnsPerHost:     options.maxConnsPerHost,
	}
	return nil
}
