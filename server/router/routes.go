package router

import (
	"net/http"
	"strings"
)

func addRoute(path, method string, handler http.Handler, prefixes ...string) {
	r := router
	if len(prefixes) > 0 {
		build := strings.Builder{}
		build.WriteRune('/')
		for i, pre := range prefixes {
			pre = strings.Trim(pre, "/")
			build.WriteString(pre)
			if i+1 < len(prefixes) {
				build.WriteRune('/')
			}
		}
		r = Subrouter(build.String())
	}
	_ = r.Handle(path, handler).Methods(method)
}

// Get adds a GET endpoint to the router.
func Get(path string, handler http.Handler, prefixes ...string) {
	addRoute(path, http.MethodGet, handler, prefixes...)
}

// Post adds a POST endpoint to the router.
func Post(path string, handler http.Handler, prefixes ...string) {
	addRoute(path, http.MethodPost, handler, prefixes...)
}

// Put adds a PUT endpoint to the router.
func Put(path string, handler http.Handler, prefixes ...string) {
	addRoute(path, http.MethodPut, handler, prefixes...)
}

// Delete adds a DELETE endpoint to the router.
func Delete(path string, handler http.Handler, prefixes ...string) {
	addRoute(path, http.MethodDelete, handler, prefixes...)
}
