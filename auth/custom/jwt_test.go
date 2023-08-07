package custom

import (
	"encoding/base64"
	"log"
	"net/url"
	"testing"
)

type confMock struct{}

func (c confMock) ClientURL() *url.URL {
	u, _ := url.Parse("https://localhost")
	return u
}

func (c confMock) ServerURL() *url.URL {
	return c.ClientURL()
}

func (c confMock) Audience() []string {
	return []string{"https://localhost"}
}

func (c confMock) Issuer() string {
	return "bentkowski.io"
}
func (c confMock) Secret(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString("+FIerIo/W71mnFDFAZ9knYWMASeps9f1QqSlJDaMbmjgN28vs4frnU1V0cfRK32e8Fu5Y1MLBf+WV/qBOdAAqjdmqI5zOLa32N2RZKPFb4glJQSAFTCD0pdEQNlcR0ARL6dNjx78oTnfsbz56v0cTJGoDou+qxdXkXIpmq8Wido=")
}

func TestJwt(t *testing.T) {
	config := confMock{}
	a := New(config)
	s1, s2, err := a.CreateClientTokens("123")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(s1)
	log.Println(s2)
}
