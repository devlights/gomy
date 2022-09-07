package chans

import "context"

// RepeatContext は、Repeat の context.Context 版です.
func RepeatContext[T any](ctx context.Context, values ...T) <-chan T {
	return Repeat(ctx.Done(), values...)
}

// Repeat -- 指定した値を永遠と繰り返すチャネルを返します。
func Repeat[T any](done <-chan struct{}, values ...T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()

	return out
}

// RepeatFn -- 指定した関数を永遠と繰り返し、その戻り値を返すチャネルを返します。
func RepeatFn[T any](done <-chan struct{}, fn func() T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case out <- fn():
			}
		}
	}()

	return out
}
