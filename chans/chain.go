package chans

import "time"

// Chain -- 指定された base チャネルがクローズした後に next で指定された関数を呼び出します.
//
// done チャネルがクローズした場合、 next は実行されません。
func Chain(done, base <-chan struct{}, next func(finishedTime time.Time)) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		select {
		case <-done:
			return
		case <-base:
			if next != nil {
				next(time.Now())
			}

			return
		}
	}()

	return ch
}
