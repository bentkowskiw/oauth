package server

import "net/url"

type Closer interface {
	Close()
}

type configer interface {
	ServerURL() *url.URL
	KeyPath() string
	CertPath() string
}
