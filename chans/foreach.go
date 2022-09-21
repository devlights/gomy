package chans

import "context"

// ForEachContext は、ForEach の context.Context 版です.
func ForEachContext[T any](ctx context.Context, in <-chan T, callback func(v T)) {
	ForEach(ctx.Done(), in, callback)
}

// ForEach -- 指定されたチャネルをループします。
func ForEach[T any](done <-chan struct{}, in <-chan T, callback func(v T)) {
	for {
		select {
		case <-done:
			return
		case v, ok := <-in:
			if !ok {
				return
			}

			callback(v)
		}
	}
}
