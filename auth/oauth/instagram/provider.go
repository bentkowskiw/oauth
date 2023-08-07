package instagram

import (
	"github.com/oauth/auth/oauth"
	"golang.org/x/oauth2"
)

func (p *provider) Name() string {
	return p.name
}
func (p *provider) OAuthConfig(b []byte) (cfg *oauth2.Config, err error) {
	return oauth.DefaultConfig(b, p)
}
func (p *provider) Endpoint() oauth2.Endpoint {
	return p.endpoint
}
