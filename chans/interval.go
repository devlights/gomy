package chans

import (
	"time"
)

// Interval -- 指定した間隔でデータを出力していくチャネルを生成します。
func Interval[T any](done <-chan struct{}, in <-chan T, interval time.Duration) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for v := range OrDone(done, in) {
			select {
			case out <- v:
				<-time.After(interval)
			case <-done:
			}
		}
	}()

	return out
}
