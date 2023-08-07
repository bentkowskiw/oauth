package restoauth

import "github.com/oauth/model"

type LoginRequired struct {
	LoginRedirectURL string   `json:"loginRedirectURL"`
	Scope            []string `json:"scope"`
}

type loginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
