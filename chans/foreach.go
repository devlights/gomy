package chans

import "context"

// ForEachContext は、ForEach の context.Context 版です.
func ForEachContext[T any](ctx context.Context, in ...T) <-chan T {
	return ForEach(ctx.Done(), in...)
}

// ForEach -- 指定されたデータを出力するチャネルを生成します。
func ForEach[T any](done <-chan struct{}, in ...T) <-chan T {
	out := make(chan T)

	go func(data []T) {
		defer close(out)

		for _, v := range data {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}(in)

	return out
}
