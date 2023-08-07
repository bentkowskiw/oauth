package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/oauth/constants"
)

// TimeOutHandler adds a max time to handle the request before timeout
func TimeOutHandler(next http.Handler, timeout time.Duration) http.Handler {
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consider timeout only if the duration is bigger than 0
		if timeout > 0 {
			ctx := r.Context()

			ctxTimeOut, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			r = r.WithContext(ctxTimeOut)
		}

		// Let's work
		next.ServeHTTP(w, r)
	})

	return f
}

func Timeout(timeout ...time.Duration) func(http.Handler) http.Handler {
	tmout := constants.RequestTimeout
	if len(timeout) == 1 {
		tmout = timeout[0]
	}
	f := func(next http.Handler) http.Handler {
		return TimeOutHandler(next, tmout)
	}
	return f
}
