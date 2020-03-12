package fileio

import (
	"io/ioutil"
	"os"
)

// TempDir は、一時ディレクトリを作成し (ディレクトリパス, defer用関数, エラー) を返します.
//
// defer用関数は、内部で os.RemoveAll(dir) を呼び出します。
func TempDir(pattern string) (string, func(), error) {
	dir, err := ioutil.TempDir("", pattern)
	if err != nil {
		return "", nil, err
	}

	removeFn := func() {
		_ = os.RemoveAll(dir)
	}

	return dir, removeFn, nil
}
