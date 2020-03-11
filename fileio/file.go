package fileio

import (
	"os"
)

// OpenAppend は、追加モードでファイルをオープンします。
func OpenAppend(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
