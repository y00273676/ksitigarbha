// Package expvars expvars represents `exported vars`
// provides api for both env and flags for variable-inputs for app configurations
// see examples for detail usage
package expvars

import (
	"errors"
	"flag"
	"strconv"
)

type DataHolder struct {
	*flag.FlagSet
	env map[string]string
}

var (
	UndefinedKeyErr = errors.New("key not defined")
)

func New(set *flag.FlagSet) *DataHolder {
	return &DataHolder{
		FlagSet: set,
		env:     map[string]string{},
	}
}

type Var interface {
	Register(holder *DataHolder)
}

// Get 只能通过 flag 的 key 来 get，不支持用 env 的 key 来做 get.
// 这样的场景是能通过 flag 和 env 两种方式来配置，但是开放的get 接口一定是 flag 的 key
func (d *DataHolder) Get(key string) (string, bool) {
	var targetFlag = d.FlagSet.Lookup(key)
	if targetFlag == nil {
		return "", false
	}
	return targetFlag.Value.String(), true
}

// EnvGet， 通过env 的 key 来 get. 这时候仅仅支持从 env 的变量里获取数据
// 应该几乎用不到. EnvGet可以拿到一个跟 flag 不一样的值，如果同时指定的话
// 但是通过 Get 获取，会有 flag > env的机制，即: flag 有值，Get 就会拿到 flag 的值
func (d *DataHolder) EnvGet(key string) (string, bool) {
	envVal, ok := d.env[key]
	return envVal, ok
}

func BoolE(key string) (bool, error) {
	strVar, ok := Get(key)
	if !ok {
		return false, UndefinedKeyErr
	}
	return strconv.ParseBool(strVar)
}
func Bool(key string) bool {
	val, _ := BoolE(key)
	return val
}

func IntE(key string) (bool, error) {
	strVar, ok := Get(key)
	if !ok {
		return false, UndefinedKeyErr
	}
	return strconv.ParseBool(strVar)
}
func Int(key string) bool {
	val, _ := BoolE(key)
	return val
}

func Int64E(key string) (int64, error) {
	strVar, ok := Get(key)
	if !ok {
		return 0, UndefinedKeyErr
	}
	return strconv.ParseInt(strVar, 10, 64)
}
func Int64(key string) int64 {
	val, _ := Int64E(key)
	return val
}

func StringE(key string) (string, error) {
	strVar, ok := Get(key)
	if !ok {
		return "", UndefinedKeyErr
	}
	return strVar, nil
}
func String(key string) string {
	val, _ := StringE(key)
	return val
}
