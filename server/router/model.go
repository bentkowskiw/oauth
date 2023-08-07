package router

import "net/http"

type scopedHandler struct {
	hFunc func(http.ResponseWriter, *http.Request)
	sFunc func() []string
}

func (h scopedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.hFunc(w, r)
}

func (h scopedHandler) Scopes() []string {
	return h.sFunc()
}
