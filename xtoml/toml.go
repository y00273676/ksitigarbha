package xtoml

import (
	"bytes"
	"fmt"

	"github.com/pelletier/go-toml"
)

// Marshal marshal a struct into toml format data.
// Be aware that when you get the marshal result,
// there will be a line break and some training spaces in the end.
// This function won't try to remove them.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, fmt.Errorf("toml.Encode: %w", err)
	}
	return buf.Bytes(), nil
}

// MarshalString will output a string instead of byte. SUGAR!
func MarshalString(v interface{}) (s string, err error) {
	var b []byte
	b, err = Marshal(v)
	if err != nil {
		return
	}
	return string(b), nil
}

// Unmarshal bytes into the given struct.
func Unmarshal(b []byte, v interface{}) error {
	return toml.NewDecoder(bytes.NewReader(b)).Decode(v)
}
