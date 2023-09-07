package oauth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func (a *oAuth) ProviderName() string {
	return a.providerName
}

// LoginSessionURL creates new oauth session.
// Returns oauth loginURL associated with session
func (a *oAuth) LoginSessionURL(ctx context.Context) (sessionID SessionUUID, url string, err error) {
	sessionID = SessionUUID(uuid.NewString())
	url = a.config.AuthCodeURL(string(sessionID), oauth2.AccessTypeOffline)
	err = a.per.Save(
		ctx,
		newOAuthSession(a.providerName, sessionID),
		a.sessionDuration,
	)
	return
}

// RequestOAuthToken checks is session exists and returns token
func (a *oAuth) RequestOAuthToken(ctx context.Context, sessionId SessionUUID, authCode string) (tokenOauth *oauth2.Token, err error) {
	if err = a.per.Read(
		ctx,
		newOAuthSession(a.providerName, sessionId),
	); err != nil {
		return
	}
	return a.config.Exchange(ctx, authCode)
}

func DefaultConfig(b []byte, p provider) (cfg *oauth2.Config, err error) {
	cf := defaultConfig{}
	if err = json.Unmarshal(b, &cf); err != nil {
		return
	}

	cfg = &oauth2.Config{
		ClientID:     cf.ClientID,
		ClientSecret: cf.ClientSecret,
		RedirectURL:  RedirectURL(p),
		Scopes:       cf.Scopes,
		Endpoint:     oauth2.Endpoint{
			AuthURL: cf.AuthURL,
			TokenURL: cf.TokenURL,
		},
	}
	return
}

func RedirectURL(p provider) string {
	return fmt.Sprintf("%s/%s/callback", p.ServerURL(), p.Name())
}
