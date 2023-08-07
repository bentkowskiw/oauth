package restoauth

import (
	"net/http"

	"github.com/oauth/server/router"
)

const prefix = "oauth"

func (h *Handler) AddRoutes() {
	router.Get("/callback/", http.HandlerFunc(h.callback), prefix, h.oa.ProviderName())
	router.Post("/callback/", http.HandlerFunc(h.callback), prefix, h.oa.ProviderName())

	router.Delete("/deauthorize/", http.HandlerFunc(h.deauthorize), prefix, h.oa.ProviderName())

	router.Get("/loginurl/", http.HandlerFunc(h.loginURL), prefix, h.oa.ProviderName())

	router.Get("/login/", http.HandlerFunc(h.loginRedirect), prefix, h.oa.ProviderName())
	router.Put("/login/", http.HandlerFunc(h.loginRedirect), prefix, h.oa.ProviderName())
	router.Post("/login/", http.HandlerFunc(h.loginRedirect), prefix, h.oa.ProviderName())

}
