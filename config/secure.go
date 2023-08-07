package config

import (
	"encoding/base64"
	"strconv"

	"github.com/oauth/lib/errlib"
)

func (s *Settings) Secure() *secure {
	return s.cfg.Secure
}

func (st *secure) CipherKey() (key [32]byte) {

	var err error
	var str string
	s, ok := st.environment.get("cipher_key", st)

	switch ok {
	case true:
		str = s
	default:
		str = st.CipherB64
	}
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	if len(b) != 32 {
		panic("invalid cipher_key length")
	}
	key = [32]byte{}
	copy(key[:], b)
	return
}
func (st *secure) Disabled() bool {
	d, ok := st.environment.get("disable", st)
	if ok {
		b, err := strconv.ParseBool(d)
		errlib.PanicOnErr(err)
		return b
	}
	return st.Disable
}

func (st *secure) prefix() []string {
	return []string{"secure"}
}
