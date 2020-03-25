package chans

import (
	"time"
)

// Interval -- 指定した間隔でデータを出力していくチャネルを生成します。
func Interval(done <-chan struct{}, in <-chan interface{}, interval time.Duration) <-chan interface{} {
	out := make(chan interface{})

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
