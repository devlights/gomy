package misc

import (
	"fmt"
	"os/exec"
	"runtime"
)

var (
	// ErrURLIsEmpty -- url が 空白だった場合のエラー
	ErrURLIsEmpty = fmt.Errorf("url is empty")
)

// OpenBrowser -- 指定されたURLでブラウザを開きます。
//
// Reference: https://github.com/go101/go101/blob/8760586e110316dcfaad36ff167f67b4f09de0a7/go101.go#L430
func OpenBrowser(url string) error {
	var (
		cmd string
		args []string
	)

	if url == "" {
		return ErrURLIsEmpty
	}

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	c := exec.Command(cmd, append(args, url)...)
	return c.Start()
}