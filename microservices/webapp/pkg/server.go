package pkg

import (
	"net/http"
	"time"
)

type ServicesTCP struct {
	AppAddr     string
	AuthAddr    string
	HistoryAddr string
	TaskAddr    string
}

type Config struct {
	ServicesTCP  ServicesTCP
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	LogFmt       string
	LogLevel     string
}

func NewServer(h http.Handler, c *Config) *http.Server {
	return &http.Server{
		Addr:         c.ServicesTCP.AppAddr,
		Handler:      h,
		IdleTimeout:  c.IdleTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
}
