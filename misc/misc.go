package misc

import (
	"context"
	"fmt"
	"math"
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
		cmd  string
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

func PrimeNumber(ctx context.Context, limit int) chan int {
	ctx, cxl := context.WithCancel(ctx)
	ch := make(chan int)

	go func() {
		defer cxl()
		defer close(ch)

		ch <- 2

	OUTER:
		for i := 3; i < limit; i += 2 {
			l := int(math.Sqrt(float64(i)))

			found := false
		INNER:
			for j := 3; j < l+1; j += 2 {
				select {
				case <-ctx.Done():
					break OUTER
				default:
					if i%j == 0 {
						found = true
						break INNER
					}
				}
			}

			if !found {
				ch <- i
			}
		}
	}()

	return ch
}
