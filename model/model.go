package model

import "golang.org/x/oauth2"

// User
type User struct {
	Uuid         string `json:"uuid"`
	ResourceUuid string `json:"resourceUuid"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Locale       string `json:"locale"`
	Phone        string `json:"phone"`
}

// UserToken
type UserToken struct {
	*User       `json:"user"`
	*OauthToken `json:"token"`
}

type OauthToken struct {
	*oauth2.Token
	scopes []string
}
