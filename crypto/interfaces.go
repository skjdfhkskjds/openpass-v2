package crypto

import (
	"github.com/skjdfhkskjds/openpass-v2/types/key"
	"github.com/skjdfhkskjds/openpass-v2/types/password"
)

// Algorithm is an interface for encryption and decryption
// of a password.
type Algorithm interface {
	Encrypt(plainText string) (string, error)
	Decrypt(password password.Password) (string, error)
}

// KeyDerivationFunction is an interface for deriving a key
// from a master password.
// This key is used as a standard length key for encryption
// and decryption by the Algorithm interface.
type KeyDerivationFunction interface {
	DeriveKey(password string, params *key.Params) ([]byte, error)
}
