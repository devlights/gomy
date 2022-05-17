package errs

import (
	"fmt"
	"io"
	"os"
)

var (
	stdOutWriter io.Writer = os.Stdout
)

// Stdout は、指定されたエラーの値を標準出力に出力して v の値だけを返します。
func Stdout[T any](v T, err error) T {
	if err != nil {
		fmt.Fprintln(stdOutWriter, err)
	}
	return v
}
