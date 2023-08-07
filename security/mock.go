package security

func NewMockCrypter() Crypter {
	return cryptoMock("")
}

func (cryptoMock) Encrypt(b []byte) []byte {
	return b
}
func (cryptoMock) Decrypt(b []byte) ([]byte, error) {
	return b, nil
}
func (cryptoMock) Hash(s string) string {
	return s
}
