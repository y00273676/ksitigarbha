package identifier_test

import (
	"encoding/json"
	"log"
	"testing"

	"ksitigarbha/identifier"
)

func TestMarshalJSON(t *testing.T) {
	var cases = []struct {
		Identifier identifier.Identifier
		JSON       string
	}{
		{
			Identifier: 123456,
			JSON:       `"123456"`,
		},
		{
			Identifier: 124151251251,
			JSON:       `"124151251251"`,
		},
	}

	for _, c := range cases {
		result, err := json.Marshal(c.Identifier)
		if err != nil {
			log.Fatalf("identifier error: %v", err)
			return
		}
		if string(result) != c.JSON {
			log.Fatalf("identifier.MarshalJSON not expected, result: %s, expect: %v", result, c.JSON)
		}
	}
}

func TestString(t *testing.T) {
	var cases = []struct {
		Identifier identifier.Identifier
		Str        string
	}{
		{
			Identifier: 123456,
			Str:        "123456",
		},
		{
			Identifier: 124151251251,
			Str:        "124151251251",
		},
	}

	for _, c := range cases {
		result := c.Identifier.String()
		if result != c.Str {
			log.Fatalf("identifier.String not expected, result: %s, expect: %v", result, c.Str)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var cases = []struct {
		JSON string
		identifier.Identifier
	}{
		{
			Identifier: 123456,
			JSON:       `"123456"`,
		},
		{
			Identifier: 124151251251,
			JSON:       `"124151251251"`,
		},
	}

	for _, c := range cases {
		var result identifier.Identifier
		var err = json.Unmarshal([]byte(c.JSON), &result)
		if err != nil {
			log.Fatalf("unmarhsal error: %v", err)
			return
		}
		if result != c.Identifier {
			log.Fatalf("identifier.UnmarshalJSON not expected, result: %s, expect: %v", result, c.Identifier)
		}
	}
}
