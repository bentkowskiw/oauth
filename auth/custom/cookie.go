package custom

import (
	"net/http"
	"time"
)

const cookieName = "authorization"

func (a *auth) CreateSessionCookie(sessionId string) (*http.Cookie, error) {

	token, err := a.CreateClientSessionToken(sessionId)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     cookieName,
		Value:    *token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 1),
		MaxAge:   10000,
		Path:     "/",
		Domain:   a.clientHost,
		// SameSite: http.SameSiteNoneMode,
		Secure: true,
	}
	return &cookie, nil
}

func (a *auth) ClearCookie(w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		Path:     "/",
		Domain:   a.clientHost,
	}
	http.SetCookie(w, &cookie)
}

func (a *auth) GetEncCookie(r *http.Request) (_ *string, err error) {

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}
	token, err := a.ValidateToken(r.Context(), cookie.Value)
	if err != nil {
		return
	}
	sess := token.Subject()
	return &sess, nil
}
