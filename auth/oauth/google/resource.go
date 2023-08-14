package google

import (
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
	serverURL string
}

func Provider(cfg configer) *provider {
	return &provider{
		name:      "google",
		endpoint:  google.Endpoint,
		serverURL: fmt.Sprint(cfg.GetServerURL()),
	}
}

type configer interface {
	GetServerURL() url.URL
}
