package middleware

import "net/http"

// ResponseWrapper implements the http.ResponseWriter interface,
// but also separately stores the response status code so it can be logged.
type ResponseWrapper struct {
	statusCode int
	err        error
	w          http.ResponseWriter
}

// Header implements the http.ResponseWriter
func (rw *ResponseWrapper) Header() http.Header {
	return rw.w.Header()
}

// Write implements the http.ResponseWriter
func (rw *ResponseWrapper) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}

// WriteHeader implements the http.ResponseWriter and saves the status code
func (rw *ResponseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.w.WriteHeader(statusCode)
}
