package password

import "fmt"

type Params struct {
	Algorithm string
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

func (p *Params) String() string {
	return fmt.Sprintf("Algorithm: %s", p.Algorithm)
}

func WithAlgorithm(algorithm string) ParamsOption {
	return func(p *Params) {
		p.Algorithm = algorithm
	}
}
