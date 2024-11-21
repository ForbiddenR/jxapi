package api

import (
	"log/slog"
	"net/url"
	"os"
	"time"

	"github.com/ForbiddenR/jxapi/v2/rest"
	"github.com/valyala/fasthttp"
)

var ServiceClient rest.Interface

var Log *slog.Logger

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
	Log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	Log.Info("InitApi", slog.String("esamUrl", esamUrl), slog.String("servicesUrl", servicesUrl))
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
	service, err := url.Parse(servicesUrl)
	if err != nil {
		return err
	}
	ServiceClient, err = rest.NewRestClient(service, rest.ClientContentConfig{}, client)
	return err
}
