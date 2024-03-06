package types

import "github.com/skjdfhkskjds/openpass/v2/types/password"

type User struct {
	Username string
	Password string

	Entries map[string]password.Password
}
