package password

import (
	"github.com/skjdfhkskjds/openpass-v2/types/key"
)

// Password is a struct that holds the encrypted password
// and the relevant parameters used to encrypt it.
type Password struct {
	Hash  []byte
	Nonce []byte

	Key    key.Key
	Params Params
}

func New(hash, nonce []byte, key key.Key, params Params) Password {
	return Password{
		Hash:   hash,
		Nonce:  nonce,
		Key:    key,
		Params: params,
	}
}
