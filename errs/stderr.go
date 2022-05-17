package errs

import (
	"fmt"
	"io"
	"os"
)

var (
	stdErrWriter io.Writer = os.Stderr
)

// Stderr は、指定されたエラーの値を標準エラー出力に出力して v の値だけを返します。
func Stderr[T any](v T, err error) T {
	if err != nil {
		fmt.Fprintln(stdErrWriter, err)
	}
	return v
}
