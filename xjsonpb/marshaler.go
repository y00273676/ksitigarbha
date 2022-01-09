package xjsonpb

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	Marshaler = jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
	}
	Unmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
)

// Marshal a proto message into bytes
func Marshal(v proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	if err := Marshaler.Marshal(&buf, v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal a proto message from bytes
func Unmarshal(b []byte, v proto.Message) error {
	return Unmarshaler.Unmarshal(bytes.NewReader(b), v)
}
