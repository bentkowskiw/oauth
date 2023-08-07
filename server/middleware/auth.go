package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/oauth/lib/ctxlib"
)

func NewAuthHandler(auth authenticator, errhandlerOpt ...errorHandler) AuthHandler {
	var errh errorHandler = DefaultAuthErrorHandler
	if len(errhandlerOpt) == 1 {
		errh = errhandlerOpt[0]
	}
	return AuthHandler{
		errorHandler:  errh,
		authenticator: auth,
	}
}

type AuthHandler struct {
	errorHandler
	authenticator
}

func (h AuthHandler) Authenticate(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		tokenString := extractToken(r)

		tok, err := h.ValidateToken(r.Context(), tokenString)
		if err != nil {
			h.HandleError(err, w, r)
			return
		}
		ctx := ctxlib.SetUserId(r.Context(), tok.Subject())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

type authenticator interface {
	ValidateToken(ctx context.Context, tokenString string) (jwt.Token, error)
}

type errorHandler interface {
	HandleError(error, http.ResponseWriter, *http.Request)
}

type ErrorHandlerAdapter func(error, http.ResponseWriter, *http.Request)

func (a ErrorHandlerAdapter) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	a(err, w, r)
}

func NewErrorHandlerAdapter(f func(error, http.ResponseWriter, *http.Request)) ErrorHandlerAdapter {
	return ErrorHandlerAdapter(f)
}
