package restoauth

import (
	"net/http"

	"github.com/oauth/lib/httplib"
)

// oauth is service authorization callback function.
// It gets called by the OAuth external service. Saves token data.[2]
func (h *Handler) callback(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")
	sessionId := r.FormValue("state")
	ctx := r.Context()
	tokenOauth, err := h.oa.RequestOAuthToken(ctx, sessionId, authCode)
	if err != nil {
		httplib.SendError(w, err)
		return
	}
	usr, err := h.usr.User(ctx, tokenOauth)
	if err != nil {
		httplib.SendError(w, err)
		return
	}

	access, refresh, err := h.c.CreateClientTokens(usr.Uuid)
	if err != nil {
		httplib.SendError(w, err)
		return
	}

	url := h.clientURL
	u := url.Scheme + "://" + url.Host

	// associate service authorization with client session ID
	// TODO: send token
	_ = access
	_ = refresh

	http.Redirect(w, r, u, http.StatusMovedPermanently)

}

// loginURL similar to login, but the client gets only url and calls external service oauth itself [1']
func (h *Handler) loginURL(w http.ResponseWriter, r *http.Request) {
	url, err := h.generateLoginURL(w, r)
	if err != nil {
		httplib.SendError(w, err)
		return
	}
	httplib.Send(w, url)
}

// login called by client to auth to external oauth [1]
// redirects user to accept access to account
func (h *Handler) loginRedirect(w http.ResponseWriter, r *http.Request) {
	url, err := h.generateLoginURL(w, r)
	if err != nil {
		httplib.SendError(w, err)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) generateLoginURL(w http.ResponseWriter, r *http.Request) (url string, err error) {
	ctx := r.Context()
	// unique authorization url associated with sessionId
	session, url, err := h.oa.LoginSessionURL(ctx)
	if err != nil {
		return
	}
	// send sessionId as cookie
	cookie, err := h.c.CreateSessionCookie(session)
	if err != nil {
		return
	}
	http.SetCookie(w, cookie)
	return url, nil
}

func (h *Handler) deauthorize(w http.ResponseWriter, r *http.Request) {

	// TODO: deauthorize by provider

}
