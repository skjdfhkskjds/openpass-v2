package chacha2

import (
	"golang.org/x/crypto/chacha20poly1305"

	"github.com/skjdfhkskjds/openpass-v2/types/key"
	"github.com/skjdfhkskjds/openpass-v2/types/password"
)

// Algorithm is an implementation of the crypto.Algorithm interface
// using the ChaCha20-Poly1305 AEAD cipher.
type Algorithm struct {
	key key.Key

	params *password.Params
}

func New(key key.Key) *Algorithm {
	return &Algorithm{
		key:    key,
		params: DefaultChaCha2Params(),
	}
}

func (a *Algorithm) Encrypt(plainText string) (password.Password, error) {
	aead, err := chacha20poly1305.NewX(a.key.Hash[:])
	if err != nil {
		return password.Password{}, err
	}

	nonce := make([]byte, aead.NonceSize())
	encryptedHash := aead.Seal(nonce, nonce, []byte(plainText), nil)

	// TODO: fill password parameters
	return password.New(encryptedHash, nonce, a.key, password.Params{}), nil
}

func (a *Algorithm) Decrypt(password password.Password) (string, error) {
	aead, err := chacha20poly1305.NewX(a.key.Hash[:])
	if err != nil {
		return "", err
	}

	decryptedPassword, err := aead.Open(nil, password.Nonce, password.Hash, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedPassword), nil
}
