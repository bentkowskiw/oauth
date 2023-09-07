package restoauth

import (
	"net/http"

	"github.com/oauth/constants"
	"github.com/oauth/server/router"
)

func (h *Handler) AddRoutes() {
	router.Get("/callback", http.HandlerFunc(h.callback), h.oa.ProviderName())
	router.Post("/callback", http.HandlerFunc(h.callback), h.oa.ProviderName())

	router.Post("/deauthorize", http.HandlerFunc(h.deauthorize), constants.ApiPrefix, h.oa.ProviderName())
	router.Get("/deauthorize", http.HandlerFunc(h.deauthorize), constants.ApiPrefix, h.oa.ProviderName())

	router.Get("/loginurl", http.HandlerFunc(h.loginURL), h.oa.ProviderName())

	router.Get("/login", http.HandlerFunc(h.loginRedirect), h.oa.ProviderName())
	router.Post("/login", http.HandlerFunc(h.loginRedirect), h.oa.ProviderName())

}
