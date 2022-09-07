package chans

import "context"

// SkipContext は、Skip の context.Context 版です.
func SkipContext[T any](ctx context.Context, in <-chan T, count int) <-chan T {
	return Skip(ctx.Done(), in, count)
}

// SkipWhileContext は、SkipWhile の context.Context 版です.
func SkipWhileContext[T comparable](ctx context.Context, in <-chan T, value T) <-chan T {
	return SkipWhile(ctx.Done(), in, value)
}

// SkipWhileFnContext は、SkipWhileFn の context.Context 版です.
func SkipWhileFnContext[T comparable](ctx context.Context, in <-chan T, fn func() T) <-chan T {
	return SkipWhileFn(ctx.Done(), in, fn)
}

// Skip -- 指定した個数分、入力用チャネルから値をスキップするチャネルを返します。
func Skip[T any](done <-chan struct{}, in <-chan T, count int) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		skipCount := 0
		for v := range OrDone(done, in) {
			if skipCount < count {
				skipCount++
				continue
			}

			select {
			case out <- v:
			case <-done:
			}
		}
	}()

	return out
}

// SkipWhile -- 入力用チャネルから取得した値が指定した値と同一である間、値をスキップし続けるチャネルを返します。
func SkipWhile[T comparable](done <-chan struct{}, in <-chan T, value T) <-chan T {
	return SkipWhileFn(done, in, func() T { return value })
}

// SkipWhileFn -- 入力用チャネルから取得した値が指定した関数の戻り値と同一である間、値をスキップし続けるチャネルを返します。
func SkipWhileFn[T comparable](done <-chan struct{}, in <-chan T, fn func() T) <-chan T {
	out := make(chan T)

	go func(fn func() T) {
		defer close(out)

		var (
			fnResult = fn()
			skipEnd  = false
		)

		for v := range OrDone(done, in) {
			if !skipEnd && v == fnResult {
				continue
			}

			skipEnd = true

			select {
			case out <- v:
			case <-done:
			}
		}
	}(fn)

	return out
}
