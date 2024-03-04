package key

import (
	"fmt"

	"github.com/skjdfhkskjds/openpass-v2/types/salt"
)

// Params is a struct that holds the parameters for the
// key derivation function
type Params struct {
	Algorithm string

	Salt       salt.Salt
	Iterations int
}

type ParamsOption func(*Params)

func NewParams(opts ...ParamsOption) *Params {
	params := DefaultParams()
	for _, opt := range opts {
		opt(params)
	}
	return params
}

// TODO: fill this with sensible default values
// maybe read from an app.toml or yaml or json file
func DefaultParams() *Params {
	return &Params{}
}

// TODO: i hate fmt
func (p *Params) String() string {
	return fmt.Sprintf("Algorithm: %s, Salt: %s, Iterations: %d", p.Algorithm, p.Salt, p.Iterations)
}

func WithAlgorithm(algorithm string) ParamsOption {
	return func(p *Params) {
		p.Algorithm = algorithm
	}
}

func WithSalt(salt salt.Salt) ParamsOption {
	return func(p *Params) {
		p.Salt = salt
	}
}

func WithIterations(iterations int) ParamsOption {
	return func(p *Params) {
		p.Iterations = iterations
	}
}
