package middleware

import (
	"net/http"

	"github.com/oauth/lib/errlib"
	"github.com/oauth/lib/httplib"
)

type DataHandler func(w http.ResponseWriter, r *http.Request) (data interface{}, er error)

func (f DataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var wrap = &ResponseWrapper{
		w: w,
	}
	var data interface{}
	var er error
	if data, er = f(w, r); er != nil {
		code, found := errlib.Code(er)
		if !found {
			code = 400
		}
		wrap.statusCode = code
		wrap.err = er

		httplib.Send(wrap, er.Error(), wrap.statusCode)
		return
	}
	httplib.Send(wrap, data, wrap.statusCode)
}
