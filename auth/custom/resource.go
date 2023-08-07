// custom authentication based on Autorization header which is exchanged betwen
// frontend and backend
package custom

import (
	"net/url"

	"github.com/oauth/constants"
)

type auth struct {
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	serverHost         url.URL
	clientHost         string
	issuer             string
	audience           []string
}

func New(con configer) *auth {

	a, err := con.Secret(constants.AccesTokenSecret)
	if err != nil {
		panic(err)
	}
	r, err := con.Secret(constants.RefreshTokenSecret)
	if err != nil {
		panic(err)
	}
	serverHost := con.ServerURL()
	clientURL := con.ClientURL()
	return &auth{
		accessTokenSecret:  a,
		refreshTokenSecret: r,
		serverHost:         *serverHost,
		clientHost:         clientURL.Hostname(),
		issuer:             con.Issuer(),
		audience:           con.Audience(),
	}
}

type configer interface {
	Secret(string) ([]byte, error)
	ServerURL() *url.URL
	ClientURL() *url.URL
	Issuer() string
	Audience() []string
}
