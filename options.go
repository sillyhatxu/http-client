package client

import (
	"github.com/sillyhatxu/retry-utils"
	"time"
)

type Config struct {
	header          map[string]string
	showResponseLog bool
	timeout         time.Duration
	attempts        uint
	delay           time.Duration
	errorCallback   retry.ErrorCallbackFunc
}

type Option func(*Config)

func Header(header map[string]string) Option {
	return func(c *Config) {
		c.header = header
	}
}

func ShowResponseLog(showResponseLog bool) Option {
	return func(c *Config) {
		c.showResponseLog = showResponseLog
	}
}

func Timeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.timeout = timeout
	}
}

func Attempts(attempts uint) Option {
	return func(c *Config) {
		c.attempts = attempts
	}
}

func Delay(delay time.Duration) Option {
	return func(c *Config) {
		c.delay = delay
	}
}

func ErrorCallback(errorCallback retry.ErrorCallbackFunc) Option {
	return func(c *Config) {
		c.errorCallback = errorCallback
	}
}
