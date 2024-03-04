package crypto

import (
	"github.com/skjdfhkskjds/openpass-v2/types/key"
)

// Algorithm is an interface for encryption and decryption
// of a password.
type Algorithm interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

// KeyDerivationFunction is an interface for deriving a key
// from a password.
// This key is used as a standard length key for encryption
// and decryption by the Algorithm interface.
type KeyDerivationFunction interface {
	DeriveKey(password string, params *key.Params) ([]byte, error)
}
