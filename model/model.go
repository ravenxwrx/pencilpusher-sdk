package model

import "net/http"

type Module interface {
	GetName() string
	GetVersion() string
	GetPrefix() string
	GetHandlers() []Handler
}

type Handler interface {
	GetPath() string
	GetMethod() string
	GetHandlerFunc() http.HandlerFunc
}
