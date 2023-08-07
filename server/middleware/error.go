package middleware

import (
	"log"
	"net/http"

	"github.com/oauth/lib/httplib"
)

type AuthErrorHandler func(error, http.ResponseWriter, *http.Request)

func (h AuthErrorHandler) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	h(err, w, r)
}

var DefaultAuthErrorHandler = AuthErrorHandler(func(err error, w http.ResponseWriter, _ *http.Request) {
	if err == nil {
		return
	}
	log.Println(err)
	httplib.Unauthorized(w)
})
