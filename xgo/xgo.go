package xgo

//SafeFunc 同一个 goroutine 执行，但是会处理内部可能的 panic
// 返回值 error 为 panic 背后的信息封装为一个 error.如果外部不关心直接 omit 即可
func SafeFunc(fn func()) error {
	return tryFunc(fn, nil)
}

// 在当前 goroutine 里，执行 fn，如果有 panic 会自动捕获
func SafeFuncWithErr(fn func() error) {
	tryFuncWithErr(fn, nil)
}

// 在当前 goroutine 里，顺序执行fns，直到有 error 产生
func SerialUntilError(fns ...func() error) func() error {
	return func() error {
		for _, fn := range fns {
			if err := tryFuncWithErr(fn, nil); err != nil {
				return err
				// return errors.Wrap(err, xstring.FunctionName(fn))
			}
		}
		return nil
	}
}

//RecoveryGoFunc 新起一个 goroutine 执行，但是会处理内部可能的 panic
// 返回值 error 为 panic 背后的信息封装为一个 error.如果外部不关心直接 omit 即可
func SafeGoFunc(fn func(), recoveryFn func(error)) {
	go func() {
		var err = tryFunc(fn, nil)
		if err != nil && recoveryFn != nil {
			recoveryFn(err)
		}
	}()
}

func SafeGoFuncWithErr(fn func() error, recoveryFn func(error)) {
	go func() {
		var err = tryFuncWithErr(fn, nil)
		if err != nil && recoveryFn != nil {
			recoveryFn(err)
		}
	}()
}
