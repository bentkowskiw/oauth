package google

import (
	"encoding/json"

	"github.com/oauth/auth/oauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (p *provider) Name() string {
	return p.name
}
func (p *provider) OAuthConfig(b []byte) (cfg *oauth2.Config, err error) {
	cr := cred{}
	if err = json.Unmarshal(b, &cr); err != nil {
		return
	}
	cr.RedirectURIs = []string{oauth.RedirectURL(p)}
	b, err = json.Marshal(config{
		Credentials: cr,
	})
	if err != nil {
		return
	}
	return google.ConfigFromJSON(b, cr.Scopes...)
}
func (p *provider) Endpoint() oauth2.Endpoint {
	return p.endpoint
}
func (p *provider) ServerURL() string {
	return p.serverURL
}
