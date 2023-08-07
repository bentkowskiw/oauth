package oauth

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func (a *oAuth) ProviderName() string {
	return a.providerName
}

// LoginSessionURL creates new oauth session.
// Returns oauth loginURL associated with session
func (a *oAuth) LoginSessionURL(ctx context.Context) (sessionID, url string, err error) {
	sessionID = uuid.NewString()
	url = a.config.AuthCodeURL(sessionID, oauth2.AccessTypeOffline)
	err = a.per.Save(
		ctx,
		newOAuthSession(a.providerName, sessionID),
		a.sessionDuration,
	)
	return
}

// RequestOAuthToken checks is session exists and returns token
func (a *oAuth) RequestOAuthToken(ctx context.Context, sessionId, authCode string) (tokenOauth *oauth2.Token, err error) {
	if err = a.per.Read(
		ctx,
		newOAuthSession(a.providerName, sessionId),
	); err != nil {
		return
	}
	return a.config.Exchange(ctx, authCode)
}

func DefaultConfig(b []byte, e endpointer) (cfg *oauth2.Config, err error) {
	cf := defaultConfig{}
	if err = json.Unmarshal(b, &cf); err != nil {
		return
	}
	cfg = &oauth2.Config{
		ClientID:     cf.ClientID,
		ClientSecret: cf.ClientSecret,
		RedirectURL:  cf.RedirectURL,
		Scopes:       cf.Scopes,
		Endpoint:     e.Endpoint(),
	}
	return
}
