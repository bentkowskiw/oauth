package custom

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/twinj/uuid"
)

// CreateClientTokens create token and refresh token in JWT to pass to the client to authenticate
func (a *auth) CreateClientTokens(serviceUserId string) (*string, *string, error) {
	now := time.Now()
	td := &TokenDetails{
		AccessToken: Token{
			Subject:  serviceUserId,
			UUID:     uuid.NewV4().String(),
			Issuer:   a.issuer,
			Issued:   now,
			Expires:  now.Add(time.Minute * 15),
			Audience: a.audience,
		},
		RefreshToken: Token{
			Subject:  serviceUserId,
			UUID:     uuid.NewV4().String(),
			Issuer:   a.issuer,
			Issued:   now,
			Expires:  now.Add(time.Hour * 24 * 7),
			Audience: a.audience,
		},
	}

	if err := prepareToken(&td.AccessToken, a.accessTokenSecret); err != nil {
		return nil, nil, err
	}
	if err := prepareToken(&td.RefreshToken, a.refreshTokenSecret); err != nil {
		return nil, nil, err
	}

	return &td.AccessToken.Value, &td.RefreshToken.Value, nil
}

func (a *auth) CreateClientSessionToken(sessionId string) (_ *string, err error) {
	//Creating tmp Token
	now := time.Now().Local()
	tok := Token{
		UUID:     sessionId,
		Subject:  sessionId,
		Issuer:   a.issuer,
		Audience: []string{a.issuer},
		Issued:   now,
		Expires:  now.Add(time.Second * 30),
	}
	if err := prepareToken(&tok, a.accessTokenSecret); err != nil {
		return nil, err
	}
	return &tok.Value, nil
}

func (a *auth) verifyToken(tokenString string) (*jwt.Token, error) {
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	//Make sure that the token method conform to "SigningMethodHMAC"
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return a.accessTokenSecret, nil
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return token, nil
	return nil, nil
}

func (a *auth) ValidateToken(ctx context.Context, tokenString string) (jwt.Token, error) {
	return jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.HS256,a.accessTokenSecret))
}

func prepareToken(tok *Token, secret []byte) (err error) {
	key, err := jwk.FromRaw(secret)
	if err != nil {
		return
	}
	//Creating Access Token
	token, err := jwt.NewBuilder().
		Expiration(tok.Expires).
		IssuedAt(tok.Issued).
		Subject(tok.Subject).
		Issuer(tok.Issuer).
		Audience(tok.Audience).
		JwtID(tok.UUID).
		Build()
	if err != nil {
		return
	}
	byt, err := json.Marshal(token)
	if err != nil {
		return
	}
	buf, err := jws.Sign(byt, jws.WithKey(jwa.HS256, key))
	if err != nil {
		return
	}
	tok.Value = string(buf)
	return nil
}
