package restoauth

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/oauth/data"
	"github.com/oauth/model"
	"golang.org/x/oauth2"
)

type Handler struct {
	c         customAuther
	usr       userer
	oa        oauther
	pers      persister
	clientURL url.URL
}

func NewHandler(usr userer, pers persister, oa oauther, c customAuther, conf configer) *Handler {

	h := &Handler{
		usr:       usr,
		pers:      pers,
		oa:        oa,
		c:         c,
		clientURL: *conf.ClientURL(),
	}
	return h
}

type oauther interface {
	RequestOAuthToken(_ context.Context, sessionId, authCode string) (*oauth2.Token, error)
	LoginSessionURL(context.Context) (string, string, error)
	ProviderName() string
}

type configer interface {
	ClientURL() *url.URL
}

type userer interface {
	User(ctx context.Context, token *oauth2.Token) (model.User, error)
}

type persister interface {
	Read(context.Context, data.Readabler) error
	Save(context.Context, data.Storabler, time.Duration) error
}

type customAuther interface {
	CreateSessionCookie(sessionId string) (*http.Cookie, error)
	CreateClientTokens(serviceUserId string) (*string, *string, error)
}
