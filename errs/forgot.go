package errs

// Forgot は、指定されたエラーの値を捨てて v の値だけを返します。
func Forgot[T any](v T, _ error) T {
	return v
}
