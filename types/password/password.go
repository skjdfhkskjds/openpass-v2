package password

import (
	"github.com/skjdfhkskjds/openpass-v2/types/key"
)

type Password struct {
	Hash []byte

	Key    key.Key
	Params Params
}

func New(hash []byte, key key.Key, params Params) Password {
	return Password{
		Hash:   hash,
		Key:    key,
		Params: params,
	}
}
