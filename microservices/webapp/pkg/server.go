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

type Logging struct {
	Format   string
	Level    string
	FilePath string
}

type Config struct {
	ServicesTCP  ServicesTCP
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Log          Logging
}

func NewServer(h http.Handler, c Config) *http.Server {
	return &http.Server{
		Addr:         c.ServicesTCP.AppAddr,
		Handler:      h,
		IdleTimeout:  c.IdleTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
}
