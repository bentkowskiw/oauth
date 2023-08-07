package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/oauth/lib/httplib"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		httplib.SendError(w, err)
		return
	}
	resp := echoResponse{
		Body:   string(b),
		Header: r.Header,
	}
	httplib.Send(w, resp, http.StatusOK)
	b, _ = json.Marshal(resp)
	log.Default().Println(string(b))
}

type echoResponse struct {
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}
