package expvars

import (
	"os"
	"strconv"
)

type BoolVar struct {
	Flag    string
	Env     string
	Default bool
	Usage   string
}

func (b *BoolVar) Register(holder *DataHolder) {
	// do not register it again unless undefined
	if holder.FlagSet.Lookup(b.Flag) == nil {
		holder.FlagSet.Bool(b.Flag, b.Default, b.Usage)
	}

	// if using os.Getenv, we can't tell if the env var is
	// set empty by intention or not set at all.
	envVal, ok := os.LookupEnv(b.Env)
	if !ok {
		return
	}
	if _, err := strconv.ParseBool(envVal); err == nil {
		theFlag := holder.FlagSet.Lookup(b.Flag)
		theFlag.DefValue = envVal
		theFlag.Value.Set(envVal)
		holder.env[b.Env] = envVal
	}
}

type IntVar struct {
	Flag    string
	Env     string
	Default int
	Usage   string
}

func (b *IntVar) Register(holder *DataHolder) {
	// do not register it again unless undefined
	if holder.FlagSet.Lookup(b.Flag) == nil {
		holder.FlagSet.Int(b.Flag, b.Default, b.Usage)
	}

	// if using os.Getenv, we can't tell if the env var is
	// set empty by intention or not set at all.
	envVal, ok := os.LookupEnv(b.Env)
	if !ok {
		return
	}
	if _, err := strconv.ParseInt(envVal, 10, 64); err == nil {
		theFlag := holder.FlagSet.Lookup(b.Flag)
		theFlag.DefValue = envVal
		theFlag.Value.Set(envVal)
		holder.env[b.Env] = envVal
	}
}

type Int64Var struct {
	Flag    string
	Env     string
	Default int64
	Usage   string
}

func (b *Int64Var) Register(holder *DataHolder) {
	// do not register it again unless undefined
	if holder.FlagSet.Lookup(b.Flag) == nil {
		holder.FlagSet.Int64(b.Flag, b.Default, b.Usage)
	}
	// if using os.Getenv, we can't tell if the env var is
	// set empty by intention or not set at all.
	envVal, ok := os.LookupEnv(b.Env)
	if !ok {
		return
	}
	if _, err := strconv.ParseInt(envVal, 10, 64); err == nil {
		theFlag := holder.FlagSet.Lookup(b.Flag)
		theFlag.DefValue = envVal
		theFlag.Value.Set(envVal)
		holder.env[b.Env] = envVal
	}
}

type StringVar struct {
	Flag    string
	Env     string
	Default string
	Usage   string
}

func (b *StringVar) Register(holder *DataHolder) {
	// do not register it again unless undefined
	if holder.FlagSet.Lookup(b.Flag) == nil {
		holder.FlagSet.String(b.Flag, b.Default, b.Usage)
	}
	// if using os.Getenv, we can't tell if the env var is
	// set empty by intention or not set at all.
	envVal, ok := os.LookupEnv(b.Env)
	if !ok {
		return
	}
	theFlag := holder.FlagSet.Lookup(b.Flag)
	theFlag.DefValue = envVal
	theFlag.Value.Set(envVal)
	holder.env[b.Env] = envVal
}
