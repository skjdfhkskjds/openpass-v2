package chacha2

import "github.com/skjdfhkskjds/openpass-v2/types/password"

func DefaultChaCha2Params() *password.Params {
	return password.NewParams(
		password.WithAlgorithm(algorithmName),
	)
}
