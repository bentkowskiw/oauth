package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type provider struct {
	name     string
	endpoint oauth2.Endpoint
	*oauth2.Config
}

func Provider() *provider {
	return &provider{
		name:     "google",
		endpoint: google.Endpoint,
	}
}
