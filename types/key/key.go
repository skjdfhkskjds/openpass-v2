package key

type Key struct {
	Hash [KeySize]byte

	Params *Params
}
