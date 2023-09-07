package facebook

import (
	"golang.org/x/oauth2"
)

func (p *provider) Name() string {
	return p.name
}
func (p *provider) OAuthConfig() (cfg *oauth2.Config) {
	return p.Config
}
func (p *provider) Endpoint() oauth2.Endpoint {
	return p.endpoint
}

func (p *provider) ServerURL() string {
	return p.serverURL
}
