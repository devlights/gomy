package errs

// Panic は、指定されたエラーの値がnilではない場合にpanicさせます。エラーの値がnilの場合はvを返します。
func Panic[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
