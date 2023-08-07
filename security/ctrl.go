// Package security is responsible for encrypt / decrypt data while storing
package security

import (
	"crypto/sha512"
	"errors"
	"fmt"
)

func (c *crypto) Encrypt(b []byte) []byte {
	return c.cipher.Seal(c.nonce, c.nonce, b, nil)
}

func (c *crypto) Decrypt(b []byte) ([]byte, error) {
	nonceSize := c.cipher.NonceSize()
	if len(b) < nonceSize {
		return nil, errors.New("error ocured")
	}

	nonce, ciphertext := b[:nonceSize], b[nonceSize:]
	plaintext, err := c.cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("error ocured")
	}
	return plaintext, nil
}

func (c *crypto) Hash(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}
