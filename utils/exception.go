package utils

// Try 捕捉异常
func Try(userFn func(), catchFn func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFn(err)
		}
	}()
	userFn()
}
