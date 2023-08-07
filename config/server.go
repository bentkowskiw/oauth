package config

import "net/url"

func (s *Settings) Server() *server {
	return s.cfg.Server
}

func (s *server) prefix() []string {
	return []string{"server"}

}

func (s *server) GetServerURL() url.URL {
	u, err := url.Parse(s.URL)
	if err != nil {
		panic(err)
	}
	return *u
}

func (s *server) GetCertPath() string {
	return s.SslCert
}

func (s *server) GetKeyPath() string {
	return s.SslKey
}

func (s *client) prefix() []string {
	return []string{"client"}

}
