// [START calendar_quickstart]
package main

import (
	"embed"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oauth/auth/custom"
	"github.com/oauth/auth/oauth"
	"github.com/oauth/auth/oauth/facebook"
	"github.com/oauth/auth/oauth/google"
	"github.com/oauth/auth/oauth/instagram"
	"github.com/oauth/config"
	"github.com/oauth/constants"
	"github.com/oauth/data/cache"
	"github.com/oauth/lib/errlib"
	"github.com/oauth/restoauth"
	"github.com/oauth/security"
	"github.com/oauth/server"
	"github.com/oauth/server/middleware"
	"github.com/oauth/server/router"
)

var (
	//go:embed cfg
	configFiles embed.FS
)

func main() {

	cfg := config.NewConfig(configFiles)

	// CRYPTO
	var crypter security.Crypter
	if cfg.Secure().Disabled() {
		crypter = security.NewMockCrypter()
	} else {
		crypter = security.NewCripter(cfg.Secure())
	}

	// CACHE
	cch, err := cache.New(cfg.Redis(), crypter)
	errlib.PanicOnErr(err)

	// AUTH
	auth := custom.New(cfg.Auth())

	// initialize OAUTH for providers

	//google
	provider := oauth.New(google.Provider(cfg.OAuth()), cch, cfg.OAuth())
	restoauth.NewHandler(nil, cch, provider, auth, cfg.Auth())

	//instagram
	provider = oauth.New(instagram.Provider(cfg.OAuth()), cch, cfg.OAuth())
	restoauth.NewHandler(nil, cch, provider, auth, cfg.Auth()).AddRoutes()

	//facebook
	provider = oauth.New(facebook.Provider(cfg.OAuth()), cch, cfg.OAuth())
	restoauth.NewHandler(nil, cch, provider, auth, cfg.Auth()).AddRoutes()
	
	// endpoints protected with custom Authenticator
	initMiddleware(router.Subrouter(constants.ApiPrefix).Router, middleware.NewAuthHandler(auth).Authenticate)

	// generic handlers setup
	_ = router.Instance().Router.Path("/echo").Handler(http.HandlerFunc(middleware.EchoHandler))
	_ = router.Instance().Router.Path("/health").Handler(http.HandlerFunc(middleware.HealthHandler))
	_ = router.Instance().Router.Methods(http.MethodOptions).Handler(middleware.NewCORS(cfg).CORSHandler(router.Instance().Router))
	server.Run(cfg.Server(), router.Instance().Router, nil)

}

func initMiddleware(m *mux.Router, functions ...mux.MiddlewareFunc) {
	fs := make([]mux.MiddlewareFunc, 0, len(functions)+2)
	fs = append(fs, middleware.RecoverPanic, middleware.Timeout())
	fs = append(fs, functions...)
	m.Use(
		fs...,
	//middleware.TraceHandler,
	)
}
