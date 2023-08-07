package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// RecoverPanic allows to recover in case of panic
func RecoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				e := fmt.Errorf("panic: %+v", r)
				log.Fatal(e)

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
