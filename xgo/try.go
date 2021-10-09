package xgo

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"log"
)

// tryFuncWithErr 类似 tryFunc，只是 fn 本身就有 error 返回值
// afterFn 为 fn 执行后一定要执行的函数，无论是否发生 panic
func tryFuncWithErr(fn func() error, afterFn func()) (ret error) {
	if afterFn != nil {
		defer afterFn()
	}
	defer func() {
		if r := recover(); r != nil {
			// the skip of 2 will show the real panic line in the fn,
			// which could be verified in the test file.
			if _, ok := r.(error); ok {
				ret = r.(error)
			} else {
				ret = fmt.Errorf("%+v", r)
			}

			_, file, line, _ := runtime.Caller(2)
			log.Printf(
				"func panic line %s %d stackTrace %+v err %v", file, line, debug.Stack(), ret)

		}
	}()
	return fn()
}

// 需要一个 fn，无参数无返回值，afterFn 用于在 fn 发生后做一些首尾处理工作，可以为 nil， 无论是否发生 panic，这个 afterFn 都要执行
//返回值 error 为
// 1. 如果 panic(error)的情况，就是内部的那个 error
// 2. 如果 panic(任意非error)的情况，做一次%+v的 format 返回一个新的 error
//内部会记录日志，并返回给前端
func tryFunc(fn func(), afterFn func()) (ret error) {
	if afterFn != nil {
		defer afterFn()
	}
	defer func() {
		_, file, line, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			log.Printf("func panic line %s %d",file, line)
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()
	fn()
	return nil
}
