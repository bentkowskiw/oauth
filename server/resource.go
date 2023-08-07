package server

import "net/url"

type Closer interface {
	Close()
}

type configer interface {
	GetServerURL() url.URL
	GetKeyPath() string
	GetCertPath() string
}
