package xstrings

import (
	"testing"
)

// Benchmark functions
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand(16)
	}
}

func TestRandString(t *testing.T) {
	n := 10
	s, _ := Rand(n)
	t.Log(s)
	if len(s) != n {
		t.Errorf("Expect rand string to have length of %d. Actual string is %s", n, s)
	}
}

func TestRandNumberString(t *testing.T) {
	n := 6
	s, _ := RandNNumber(n)
	t.Log(s)
	if len(s) != n {
		t.Errorf("Expect RandNumberString result to have length of %d. Got %d", n, len(s))
	}
}
