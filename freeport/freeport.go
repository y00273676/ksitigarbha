// Package freeport let user to check any available port
// on local os and return it for further usage.
// It is very useful when you want to run some test that
// needs to occupy a real port.
package freeport

import "net"

// FreePort returns available port of the host
func FreePort() (uint16, error) {
	l, err := net.Listen("tcp", "[::]:0")
	if err != nil {
		return 0, err
	}

	p := l.Addr().(*net.TCPAddr).Port
	err = l.Close()
	return uint16(p), err
}
