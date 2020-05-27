package deepcopy

import (
	"bytes"
	"encoding/gob"
)

// GobCopy -- from に指定されたオブジェクトを to に encoding/gob を使ってコピーします。
func GobCopy(from, to interface{}) error {
	buf := new(bytes.Buffer)

	encoder := gob.NewEncoder(buf)
	decoder := gob.NewDecoder(buf)

	err := encoder.Encode(from)
	if err != nil {
		return err
	}

	err = decoder.Decode(to)
	if err != nil {
		return err
	}

	return nil
}
