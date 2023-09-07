package oauth

import (
	"context"
	"time"

	"github.com/oauth/data"
	"golang.org/x/oauth2"
)

const defaultSessionDuration = time.Second * 30

func New(provider provider, per persister, con configer) *oAuth {
	return &oAuth{
		providerName:    provider.Name(),
		config:          provider.OAuthConfig(),
		per:             per,
		sessionDuration: con.SessionDuration(defaultSessionDuration),
	}
}

type oAuth struct {
	providerName    string
	config          *oauth2.Config
	per             persister
	sessionDuration time.Duration
}

type configer interface {
	SessionDuration(defaultDur time.Duration) time.Duration
}
type persister interface {
	Read(context.Context, data.Readabler) error
	Save(context.Context, data.Storabler, time.Duration) error
}

type provider interface {
	Endpoint() oauth2.Endpoint
	Name() string
	OAuthConfig() *oauth2.Config
	ServerURL() string
}
