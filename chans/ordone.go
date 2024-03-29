package chans

import "context"

// OrDoneContext は、OrDone の context.Context 版です.
func OrDoneContext[T any](ctx context.Context, in <-chan T) <-chan T {
	return OrDone(ctx.Done(), in)
}

// OrDone -- 指定された終了チャネルと入力用チャネルのどちらかが閉じたら閉じるチャネルを返します。
func OrDone[T any](done <-chan struct{}, in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case out <- v:
				case <-done:
				}
			}
		}
	}()

	return out
}
