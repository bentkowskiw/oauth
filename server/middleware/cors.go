package middleware

import (
	"net/http"
)

type configer interface {
	Cors() []string
}

// CORSHandler middleware checks each request for it's origin and sets Access-Control headers
func (cor *Cors) CORSHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			for _, o := range cor.cons {
				if o == origin {
					w.Header().Set("Access-Control-Allow-Origin", o)
					break
				}
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, enctype, Pragma, Cache-Control")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")

			w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow cookie to be sent
		}

		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		// Lets work
		next.ServeHTTP(w, r)
	})
}

func (cor *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		for _, o := range cor.cons {
			if o == origin {
				w.Header().Set("Access-Control-Allow-Origin", o)
				break
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, enctype, Pragma, Cache-Control")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow cookie to be sent
	}

	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

}

type Cors struct {
	cons []string
}

func NewCORS(conf configer) *Cors {
	return &Cors{
		cons: conf.Cors(),
	}
}
