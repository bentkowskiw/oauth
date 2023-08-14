package google

type cred struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURIs []string `json:"redirect_uris"`
	AuthURI      string   `json:"auth_uri"`
	TokenURI     string   `json:"token_uri"`
	Scopes       []string `json:"scopes"`
}

type config struct {
	Credentials cred `json:"web"`
}
