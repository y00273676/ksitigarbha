package cmd

import (
	"os"
	"os/exec"
)

type Type uint8

const (
	TypeShell Type = iota
)

// bin/sh -c ...
func Exec(t Type, commands ...string) error {
	switch t {
	case TypeShell:
		return execShell(commands...)
	default:
		panic("not implemented")
	}
}
func ExecWithOutput(t Type, commands ...string) ([]byte, error) {

	switch t {
	case TypeShell:
		return execShellWithOutput(commands...)
	default:
		panic("not implemented")
	}
}

func execShell(commands ...string) error {
	args := append([]string{"-c"}, commands...)
	cmd := exec.Command("/bin/sh", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func execShellWithOutput(commands ...string) ([]byte, error) {
	args := append([]string{"-c"}, commands...)
	cmd := exec.Command("/bin/sh", args...)
	cmd.Env = os.Environ()
	return cmd.Output()
}
