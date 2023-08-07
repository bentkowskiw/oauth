package security

import "crypto/cipher"

type crypto struct {
	nonce  []byte
	cipher cipher.AEAD
}
type cryptoMock string
