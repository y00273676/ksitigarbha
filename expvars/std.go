package expvars

import (
	"flag"
	"os"
)

var (
	std = New(flag.CommandLine)
)

func Get(key string) (string, bool) {
	return std.Get(key)
}
func EnvGet(key string) (string, bool) {
	return std.EnvGet(key)
}

// Register 注册各种变量，巧妙利用了在注册 flag 的时候，就已经能获取的 env 的值
// 这样在注册 flag 的时候，就可以把 env 的值 set 为 flag 的默认值(如果 env 有设置值的话)
// 这样如果在 flag 没设置值的时候，因为设置了 env 的值是默认值，所以获取到的就是 env 的值了
// 如果 flag 设置了值，这时候 flag 的值会覆盖掉默认值，这时候就达到了 flag 优先
func Register(vars ...Var) {
	for _, v := range vars {
		v.Register(std)
	}
}

func Parse() error {
	return std.FlagSet.Parse(os.Args[1:])
}
