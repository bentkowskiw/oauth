package router

import (
	"github.com/gorilla/mux"
)

var (
	routers map[string]*Mux = map[string]*Mux{}
	router                  = &Mux{
		Router: mux.NewRouter().StrictSlash(true),
	}
)

// Mux wraps the gorilla router
type Mux struct {
	*mux.Router
	Prefix string
}

// New initializes and returns the Gorilla Router
func Instance() *Mux {
	return router
}

func Subrouter(prefix string) *Mux {
	r := routers[prefix]
	if r == nil {
		r = &Mux{
			Router: router.PathPrefix(prefix).Subrouter(),
			Prefix: prefix,
		}
		routers[prefix] = r
	}
	return r
}
