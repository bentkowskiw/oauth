package server

import (
	"net/http"
)

// Server to represent a http server
type Server struct {
	*http.Server
	configer
	closers []Closer
}
