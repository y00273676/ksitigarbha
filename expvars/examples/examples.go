package main

import (
	"fmt"
	"ksitigarbha/expvars"
)

// go run examples.go --watch=true --addr=localhost:8080
// export WATCH=true ADDR=localhost:8080; go run examples.go
// export WATCH=true; go run examples.go --addr=localhost:8080
// export ADDR=localhost:8080; go run examples.go --watch=true

func main() {
	var bv = &expvars.BoolVar{
		Flag:    "watch",
		Env:     "WATCH",
		Default: false,
		Usage:   "define watch changes or not",
	}
	var sv = &expvars.StringVar{
		Flag:    "addr",
		Env:     "ADDR",
		Default: "localhost:9100",
		Usage:   "define address to communicate",
	}
	expvars.Register(bv, sv)
	expvars.Parse()
	var (
		strVal  string
		boolVar bool
		ok      bool
	)
	strVal, ok = expvars.Get("addr")
	if !ok {
		fmt.Println(`Get("addr") not found`)
	}
	fmt.Printf("get by flag key[%s]:%s \n", "addr", strVal)
	strVal, ok = expvars.EnvGet("ADDR")
	if !ok {
		fmt.Println(`EnvGet("ADDR") not found`)
	}
	fmt.Printf("get by env key[%s]:%s \n", "ADDR", strVal)
	strVal = expvars.String("addr")
	fmt.Println("get by String(addr) API: ", strVal)
	strVal = expvars.String("ADDR")
	fmt.Println("get by String(ADDR) API: ", strVal)

	fmt.Println("-----------------")
	strVal, ok = expvars.Get("watch")
	if !ok {
		fmt.Println(`Get("watch") not found`)
	}
	fmt.Printf("get by flag key[%s]:%s \n", "watch", strVal)
	strVal, ok = expvars.EnvGet("WATCH")
	if !ok {
		fmt.Println(`EnvGet("WATCH") not found`)
	}
	fmt.Printf("get by env key[%s]:%s \n", "WATCH", strVal)
	boolVar = expvars.Bool("WATCH")
	fmt.Printf("get by Bool(WATCH) API[%s]:%v \n", "WATCH", boolVar)
	boolVar = expvars.Bool("watch")
	fmt.Printf("get by Bool(watch) API[%s]: %v \n", "watch", boolVar)
}
