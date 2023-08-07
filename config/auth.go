package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"time"
)

func (s *Settings) Auth() *auth {
	return s.cfg.Auth
}

func (o *auth) prefix() []string {
	return []string{"auth"}
}

func parseUrl(s string) *url.URL {
	url, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return url
}

func (o *auth) Issuer() string {
	return o.ClientURL().String()
}

func (o *auth) Audience() []string {
	return *o.cors
}

func (o *auth) ServerURL() *url.URL {
	if s, ok := o.environment.get("url", o.server); ok {
		return parseUrl(s)
	}
	return parseUrl(o.server.URL)
}

func (o *auth) ClientURL() *url.URL {
	if s, ok := o.environment.get("url", o.client); ok {
		return parseUrl(s)
	}
	return parseUrl(o.client.URL)
}

func (o *auth) LoginURLDuration() time.Duration {

	s, ok := o.environment.get("login_url_duration", o)

	if !ok {
		s = o.LoginUrlDuration
	}

	d, e := time.ParseDuration(s)
	if e != nil {
		log.Default().Printf("invalid  or missing login_url_duration")
		d = time.Second * 15
	}
	return d
}

func (o *auth) Secret(name string) ([]byte, error) {
	v, ok := o.environment.get(name, o)
	if !ok {
		switch name {
		case "access_token_secret":
			v = o.AccessTokenSecret
		case "refresh_token_secret":
			v = o.RefreshTokenSecret
		default:
			return nil, fmt.Errorf("invalid secret name: '%s'", name)
		}
	}
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}
