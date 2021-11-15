package xtoml_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "ksitigarbha/xtoml"
)

func TestMarshal(t *testing.T) {
	table := []struct {
		Input     interface{}
		Expect    string
		ExpectErr bool
	}{
		{
			Input:     1,
			ExpectErr: true,
		},
		{
			Input: struct {
				A string
			}{
				A: "a",
			},
			Expect: "A = \"a\"\n",
		},
		{
			Input: struct {
				A struct {
					B string `toml:"b"`
				} `toml:"a"`
			}{
				A: struct {
					B string `toml:"b"`
				}{
					B: "b",
				},
			},
			Expect: "\n[a]\n  b = \"b\"\n",
		},
	}

	for _, row := range table {
		b, err := Marshal(row.Input)

		if row.ExpectErr {
			if err == nil {
				t.Fatalf("expect err when marshaling input %+v, got nil", row.Input)
			}
			continue
		}

		if err != nil {
			t.Fatalf("failed to marshal input %+v, err: %s", row.Input, err)
		}

		s := string(b)
		if diff := cmp.Diff(s, row.Expect); len(diff) > 0 {
			// let's use go comp to prettify the diff
			t.Fatalf("expect marshal result to be identical, got diff:\n%s", diff)
		}
	}
}
