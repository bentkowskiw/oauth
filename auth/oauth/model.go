package oauth

type defaultConfig struct {
	// ClientID is the application's ID.
	ClientID string `json:"client_id"`
	// ClientSecret is the application's secret.
	ClientSecret string `json:"client_secret"`
	// Scope specifies optional requested permissions.
	Scopes   []string `json:"scopes"`
	AuthURL  string   `json:"auth_uri"`
	TokenURL string   `json:"token_uri"`
}

type oAuthSession struct {
	sessionData
}
type sessionData struct {
	ProviderName string      `json:"providerName"`
	SessionUUID  SessionUUID `json:"sessionUUID"`
}

type SessionUUID string
