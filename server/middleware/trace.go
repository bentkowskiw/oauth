package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TraceHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("REQUEST: %s %s%s %s %v\n\r", r.Method, r.Host, r.URL, r.Referer(), print(r.Header))
		next.ServeHTTP(w, r)
		log.Default().Printf("RESPONSE: %v \n\r", print(w.Header()))
	})
}

func print(h http.Header) string {
	j, _ := json.Marshal(h)
	return fmt.Sprintf("header: %s", j)

}
