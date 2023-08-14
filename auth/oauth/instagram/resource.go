package instagram

import (
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/instagram"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
	serverURL string
}

func Provider(cfg configer) *provider {
	return &provider{
		name:     "instagram",
		endpoint: instagram.Endpoint,
		serverURL: fmt.Sprint(cfg.GetServerURL()),
	}
}

type configer interface {
	GetServerURL() url.URL
}
