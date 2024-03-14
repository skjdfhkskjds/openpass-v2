package passwordparams

import (
	"bytes"
	"encoding/gob"
)

func NewFromBytes(b []byte) (*Params, error) {
	buffer := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buffer)
	var p Params
	if err := decoder.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Params) Bytes() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(p); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *Params) BytesUnsafe() []byte {
	bz, err := p.Bytes()
	if err != nil {
		panic(err)
	}

	return bz
}
