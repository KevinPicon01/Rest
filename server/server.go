package server

import (
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}
