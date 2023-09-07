package config

import (
	"io/fs"
	"time"
)

// Settings for the whole application
type Settings struct {
	flags       *flags
	environment *environment
	fs          fs.FS
	cfg         *cfg
	oAuth       *oAuth
	generic     map[string]interface{}
}

type cfg struct {
	Server *server `json:"server"`
	Client *client `json:"client"`
	Auth   *auth   `json:"auth"`
	Redis  *redis  `json:"redis"`
	Db     *db     `json:"db"`
	Sms    *sms    `json:"sms"`
	CORS   *cors   `json:"cors"`
	Secure *secure `json:"secure"`
}

// Server settings
type server struct {
	*environment
	URL     string `json:"url"`
	SslCert string `json:"ssl_cert"`
	SslKey  string `json:"ssl_key"`
}

type client struct {
	URL string `json:"url"`
}

type auth struct {
	cors   *cors
	client *client
	server *server
	*environment
	AccessTokenSecret  string `json:"access_token_secret"`
	RefreshTokenSecret string `json:"refresh_token_secret"`
	LoginUrlDuration   string `json:"login_url_duration"`
}

type db struct {
	*environment
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Schema   string `json:"schema"`
	Ssl      string `json:"ssl"`
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type redis struct {
	DSN string `json:"dsn"`
	*environment
}

type environment struct {
	variable map[string]string
}

type sms struct {
	Config map[string]interface{} `json:"sms"`
	*environment
}
type cors []string

type secure struct {
	CipherB64 string `json:"cipher_key"`
	Disable   bool   `json:"disable"`
	*environment
}

type flags struct {
	flags map[string]*string
}

type oAuth struct {
	oAuthProviders  map[string]any
	sessionDuration time.Duration
	*cfg
}
