package user

import "github.com/skjdfhkskjds/openpass/v2/types/key"

type Data struct {
	Username string

	Key *key.Key
}

func New(username string, key *key.Key) *Data {
	return &Data{
		Username: username,
		Key:      key,
	}
}
