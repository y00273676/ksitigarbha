package xslice_test

import (
	"testing"

	"ksitigarbha/xslice"
)

func TestIdInSlice(t *testing.T) {
	s := []int64{1, 2, 3, 4, 5, 6}

	if xslice.Int.Contains(s, 666) {
		t.Error("Expect 666 not in slice")
	}
	if !xslice.Int.Contains(s, 1) {
		t.Error("Expect 1 in slice")
	}
}

func TestIdInSlicePos(t *testing.T) {
	s := []int64{3, 5, 7, 11, 13}
	for i, v := range s {
		pos := xslice.Int64.Lookup(s, v)
		if i != pos {
			t.Errorf("Expect %d to have pos of %d. Actual %d", v, pos, i)
		}
	}

	pos := xslice.Int64.Lookup(s, 666)
	if pos != -1 {
		t.Error("Expect pos to be -1")
	}

}

func TestStrInSlice(t *testing.T) {
	s := []string{"something", "really", "bad"}
	pos := xslice.String.Lookup(s, "no thanks")
	if pos >= 0 {
		t.Error("Expect sub string no to be found.")
	}

	pos = xslice.String.Lookup(s, "bad")
	if pos < 0 {
		t.Error("Expect sub string to be found.")
	}
}
