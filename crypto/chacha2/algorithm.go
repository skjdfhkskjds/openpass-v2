package chacha2

import (
	"golang.org/x/crypto/chacha20poly1305"

	"github.com/skjdfhkskjds/openpass-v2/types/key"
	"github.com/skjdfhkskjds/openpass-v2/types/password"
)

type Algorithm struct {
	key key.Key
}

func New(key key.Key) *Algorithm {
	return &Algorithm{
		key: key,
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
	return password.New(encryptedHash, a.key, password.Params{}), nil
}
