package xerrors_test

import (
	"errors"
	"strings"
	"testing"

	. "ksitigarbha/xerrors"
)

func TestNewErrors(t *testing.T) {
	table := []struct {
		newArgs        []error
		appendErrs     []error
		expectLen      int
		expectValid    bool
		expectKeywords []string
	}{
		{
			expectLen:   0,
			expectValid: false,
		},
		{
			newArgs:        []error{errors.New("1"), errors.New("2")},
			expectLen:      2,
			expectValid:    true,
			expectKeywords: []string{"1", ";", "2"},
		},
		{
			newArgs:        []error{errors.New("1"), errors.New("2")},
			appendErrs:     []error{errors.New("3"), errors.New("4")},
			expectLen:      4,
			expectValid:    true,
			expectKeywords: []string{"1", ";", "2", "3", "4"},
		},
	}

	for _, row := range table {
		errs := NewErrors(row.newArgs...)
		errs.Append(row.appendErrs...)

		if errs.Len() != row.expectLen {
			t.Fatalf("expect len to be %d, got %d", row.expectLen, errs.Len())
		}
		if errs.Valid() != row.expectValid {
			t.Fatalf("expect valid to be %v, got %v", row.expectValid, errs.Valid())
		}

		s := errs.Error()
		t.Logf("err output: %s", s)
		for _, k := range row.expectKeywords {
			if !strings.Contains(s, k) {
				t.Fatalf("expect to have keyword %s, got %s", k, s)
			}
		}
	}
}
