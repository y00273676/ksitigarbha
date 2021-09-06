package cmd_test

import (
	"go.planetmeican.com/nerds/bazaar/cmd"
	"log"
	"testing"
)

func TestExec(t *testing.T) {
	err := cmd.Exec(cmd.TypeShell, "echo hi")
	if err != nil {
		log.Fatalf("error happened")
	}
}
func TestExecWithOutput(t *testing.T) {
	output, err := cmd.ExecWithOutput(cmd.TypeShell, "ls -h|grep test")
	if err != nil {
		log.Fatalf("error happened")
	}
	if string(output) != "cmd_test.go\n" {
		log.Fatalf("wrong expectation,output:`%s`", string(output))
	}
}
