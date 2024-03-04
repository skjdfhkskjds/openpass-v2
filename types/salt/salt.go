package salt

import "crypto/rand"

// A Salt is a random value used to increase the security of
// a hash.
// It helps prevent precomputed rainbow table attacks.
// Since it is unique for each password, we must store the
// generated salt as a parameter of the password in the database.
type Salt []byte

// New creates a new salt with size Length
func New() (Salt, error) {
	salt := make(Salt, Length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
