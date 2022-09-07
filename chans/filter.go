package chans

import "context"

// FilterContext は、Filter の context.Context 版です.
func FilterContext[T any](ctx context.Context, in <-chan T, predicate func(T) bool) <-chan T {
	return Filter(ctx.Done(), in, predicate)
}

// Filter -- 入力データチャネル in から取得したデータを predicate に渡して 真(true) となったデータを返すチャネルを生成します。
func Filter[T any](done <-chan struct{}, in <-chan T, predicate func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

	ChLoop:
		for {
			select {
			case <-done:
				break ChLoop
			case v, ok := <-in:
				if !ok {
					break ChLoop
				}

				if predicate(v) {
					select {
					case out <- v:
					case <-done:
					}
				}
			}
		}
	}()

	return out
}
