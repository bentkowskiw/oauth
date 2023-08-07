package oauth

import (
	"context"
	"time"

	"github.com/oauth/data"
	"github.com/oauth/lib/errlib"
	"golang.org/x/oauth2"
)

const defaultSessionDuration = time.Second * 30

func New(provider provider, per persister, con configer) *oAuth {
	b, err := con.ConfigData(provider.Name())
	errlib.PanicOnErr(err)
	config, err := provider.OAuthConfig(b)
	errlib.PanicOnErr(err)
	return &oAuth{
		providerName:    provider.Name(),
		config:          config,
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
	ConfigData(string) ([]byte, error)
}
type persister interface {
	Read(context.Context, data.Readabler) error
	Save(context.Context, data.Storabler, time.Duration) error
}

type provider interface {
	Name() string
	OAuthConfig([]byte) (*oauth2.Config, error)
}

type endpointer interface {
	Endpoint() oauth2.Endpoint
}
