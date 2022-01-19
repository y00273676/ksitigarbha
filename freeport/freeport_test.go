package freeport_test

import (
	"log"
	"testing"

	. "ksitigarbha/freeport"
)

func TestFreePort(t *testing.T) {
	p, err := FreePort()
	if err != nil {
		log.Fatal(err)
	}
	if p == 0 {
		log.Fatalf("expect free port to be greater than 0")
	}
	t.Logf("free port returns %d", p)
}
