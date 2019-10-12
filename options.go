package client

import (
	"time"
)

type Config struct {
	header          map[string]string
	showResponseLog bool
	timeout         time.Duration
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
