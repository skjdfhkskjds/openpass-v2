package user

import (
	"bytes"
	"encoding/gob"
)

// NewFromBytes returns a new user.Data struct from bytes
func NewFromBytes(b []byte) (*Data, error) {
	buffer := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buffer)
	var d Data
	if err := decoder.Decode(&d); err != nil {
		return nil, err
	}

	return &d, nil
}

// Bytes returns the bytes representation of a user.Data struct
func (d *Data) Bytes() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(d); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (d *Data) BytesUnsafe() []byte {
	bz, err := d.Bytes()
	if err != nil {
		panic(err)
	}

	return bz
}
