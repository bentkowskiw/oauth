package oauth

type defaultConfig struct {
	// ClientID is the application's ID.
	ClientID string `json:"client_id"`

	// ClientSecret is the application's secret.
	ClientSecret string `json:"client_secret"`

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string `json:"redirect_url"`

	// Scope specifies optional requested permissions.
	Scopes []string `json:"scopes"`
}

type oAuthSession struct {
	sessionData
}
type sessionData struct {
	ProviderName string `json:"providerName"`
	SessionUUID  string `json:"sessionUUID"`
}
