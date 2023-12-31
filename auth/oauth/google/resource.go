package google

import (
	"encoding/json"
	"net/url"

	"github.com/oauth/auth/oauth"
	"github.com/oauth/lib/errlib"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
	serverURL string
}

func Provider(cfg configer) (p *provider) {
	cr := cred{}
	name := "google"
	p = &provider{
		name:      name,
		endpoint:  google.Endpoint,
		serverURL: cfg.ServerURL().String(),
	}
	b, err := cfg.ConfigData(name)
	errlib.PanicOnErr(err)
	errlib.PanicOnErr(
		json.Unmarshal(b, &cr),
	)
	cr.RedirectURIs = []string{oauth.RedirectURL(p)}
	b, err = json.Marshal(config{
		Credentials: cr,
	})
	errlib.PanicOnErr(err)
	p.Config, err = google.ConfigFromJSON(b, cr.Scopes...)
	errlib.PanicOnErr(err)
	return p
}

type configer interface {
	ServerURL() *url.URL
	ConfigData(string) ([]byte, error)
}
