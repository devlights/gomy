package chans

// Take -- 指定した個数分、入力用チャネルから値を取得するチャネルを返します。
func Take[T any](done <-chan struct{}, in <-chan T, count int) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for i := 0; i < count; i++ {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					break
				}

				select {
				case <-done:
				case out <- v:
				}
			}
		}
	}()

	return out
}

// TakeWhile -- 入力用チャネルから取得した値が指定した値と同一である間、値を取得し続けるチャネルを返します。
func TakeWhile[T comparable](done <-chan struct{}, in <-chan T, value T) <-chan T {
	return TakeWhileFn(done, in, func() T { return value })
}

// TakeWhileFn -- 入力用チャネルから取得した値が指定した関数の戻り値と同一である間、値を取得し続けるチャネルを返します。
func TakeWhileFn[T comparable](done <-chan struct{}, in <-chan T, fn func() T) <-chan T {
	out := make(chan T)

	go func(fn func() T) {
		defer close(out)

		r := fn()
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				if v != r {
					return
				}

				select {
				case out <- v:
				case <-done:
				}
			}
		}
	}(fn)

	return out
}
