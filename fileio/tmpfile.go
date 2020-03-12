package fileio

import (
	"fmt"
	"path/filepath"

	"github.com/rs/xid"
)

// TempFileName は、一時ファイルのパスを生成して返します。
func TempFileName(dir string, prefix string) string {
	id := xid.New()
	name := fmt.Sprintf("%s-%s", prefix, id.String())
	fpath := filepath.Join(dir, name)

	return fpath
}
