package chans

import "context"

// GeneratorContext は、Generator の context.Context 版です.
func GeneratorContext[T any](ctx context.Context, in ...T) <-chan T {
	return Generator(ctx.Done(), in...)
}

// Generator -- 指定されたデータを出力するチャネルを生成します。
func Generator[T any](done <-chan struct{}, in ...T) <-chan T {
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
