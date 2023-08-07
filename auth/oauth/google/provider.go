package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (p *provider) Name() string {
	return p.name
}
func (p *provider) OAuthConfig(b []byte) (cfg *oauth2.Config, err error) {
	return google.ConfigFromJSON(b)
}
func (p *provider) Endpoint() oauth2.Endpoint {
	return p.endpoint
}
