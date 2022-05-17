package errs

// Forget は、指定されたエラーの値を捨てて v の値だけを返します。
func Forget[T any](v T, _ error) T {
	return v
}

// Drop は、Forget のエイリアスです。
func Drop[T any](v T, err error) T {
	return Forget(v, err)
}
