package instagram

import (
	"fmt"
	"net/url"

	"github.com/oauth/auth/oauth"
	"github.com/oauth/lib/errlib"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/instagram"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
	serverURL string
}

func Provider(cfg configer) (p *provider) {
	name := "instagram"
	p = &provider{
		name:      name,
		endpoint:  instagram.Endpoint,
		serverURL: fmt.Sprint(cfg.ServerURL()),
	}
	b, err := cfg.ConfigData(name)
	errlib.PanicOnErr(err)

	p.Config, err = oauth.DefaultConfig(b, p)
	errlib.PanicOnErr(err)

	return
}

type configer interface {
	ConfigData(string) ([]byte, error)
	ServerURL() *url.URL
}
