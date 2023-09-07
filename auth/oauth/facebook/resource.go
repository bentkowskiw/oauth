package facebook

import (
	"fmt"
	"net/url"

	"github.com/oauth/auth/oauth"
	"github.com/oauth/lib/errlib"
	"golang.org/x/oauth2"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
	serverURL string
}

func Provider(cfg configer) (p *provider) {
	name := "facebook"
	p = &provider{
		serverURL: fmt.Sprint(cfg.ServerURL()),
		name:      name,
	}
	b, err := cfg.ConfigData(name)
	errlib.PanicOnErr(err)

	p.Config, err = oauth.DefaultConfig(b, p)
	errlib.PanicOnErr(err)

	return
}

type configer interface {
	ServerURL() *url.URL
	ConfigData(string) ([]byte, error)
}
