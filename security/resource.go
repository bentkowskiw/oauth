package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func NewCripter(cfg configer) *crypto {
	arr := cfg.CipherKey()
	c, e := aes.NewCipher(arr[:])
	if e != nil {
		panic(e)
	}
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		panic(err)
	}
	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	return &crypto{
		cipher: gcm,
		nonce:  nonce,
	}
}

type configer interface {
	CipherKey() [32]byte
}

type Crypter interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) ([]byte, error)
	Hash(string) string
}
