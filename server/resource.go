package server

import "net/url"

type Closer interface {
	Close()
}

type configer interface {
	BindURL() *url.URL
	KeyPath() string
	CertPath() string
}
