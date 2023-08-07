package httplib

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/oauth/lib/errlib"
)

const (
	JSON = "application/json"
)

func BodyAsStruct(r *http.Request, str interface{}) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, str)
}

// SetHeaders sets the response with the default http headers (security)
func SetHeaders(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Add("Expires", "0")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("X-Frame-Options", "SAMEORIGIN")
	w.Header().Add("X-Xss-Protection", "1; mode=block")
	w.Header().Add("X-Content-Type-Options", "nosniff")
}

func MainStatus(status int) int {
	status = status / 100
	return status
}

func SendRaw(w http.ResponseWriter, content interface{}, statusOpt ...int) {
	status := http.StatusOK
	if len(statusOpt) == 1 {
		status = statusOpt[0]
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(content)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func Send(w http.ResponseWriter, content interface{}, statusOpt ...int) {
	status := http.StatusOK
	if len(statusOpt) == 1 {
		status = statusOpt[0]
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = content
	resp["status"] = status

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func SendError(w http.ResponseWriter, e error) {
	status := http.StatusInternalServerError
	msg := e.Error()
	if st, ok := errlib.Code(e); ok {
		status = st
	}
	Send(w, msg, status)
}

func SendOK(w http.ResponseWriter) {
	Send(w, "ok", http.StatusOK)
}

func Unauthorized(w http.ResponseWriter) {
	SendError(w, errlib.WithCode(errors.New("unauthorized"), http.StatusUnauthorized))
}

func GetParam(r *http.Request, key string) *string {
	keys, ok := r.URL.Query()[key]
	if !ok {
		return nil
	}
	if len(keys) == 0 {
		return nil
	}
	return &keys[0]
}
