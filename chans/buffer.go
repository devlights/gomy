package chans

import "context"

// BufferContext は、Bridge の context.Context 版です.
func BufferContext[T any](ctx context.Context, in <-chan T, count int) <-chan []T {
	return Buffer(ctx.Done(), in, count)
}

// Buffer は、入力を指定した件数分に束ねてデータを返すチャネルを生成します.
func Buffer[T any](done <-chan struct{}, in <-chan T, count int) <-chan []T {
	out := make(chan []T)

	go func() {
		defer close(out)

		for {
			items := make([]T, 0, count)
			for item := range Take(done, in, count) {
				items = append(items, item)
			}

			if len(items) == 0 {
				break
			}

			select {
			case <-done:
			case out <- items:
			}

			if len(items) != count {
				break
			}
		}
	}()

	return out
}
