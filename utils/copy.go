package utils

import (
	"bytes"
	"encoding/gob"
)

func DeepCopyGob(src, dst interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err := enc.Encode(src); err != nil {
		return err
	}

	return dec.Decode(dst)
}
