package middleware

import (
	"net/http"

	"github.com/oauth/lib/httplib"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	httplib.SendOK(w)
}
