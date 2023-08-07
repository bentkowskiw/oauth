package custom

import "time"

type TokenDetails struct {
	AccessToken  Token `json:"accessToken"`
	RefreshToken Token `json:"refreshToken"`
}

type Token struct {
	Subject  string    `json:"sub"`
	UUID     string    `json:"jti"`
	Issuer   string    `json:"iss"`
	Issued   time.Time `json:"iat"`
	Expires  time.Time `json:"exp"`
	Audience []string  `json:"aud"`
	Value    string    `json:"value"`
}

type AccessDetails struct {
	AccessTokenUid string
	UserId         string
}
