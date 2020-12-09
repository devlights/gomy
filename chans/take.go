package chans

// Take -- 指定した個数分、入力用チャネルから値を取得するチャネルを返します。
func Take(done <-chan struct{}, in <-chan interface{}, count int) <-chan interface{} {
	out := make(chan interface{})

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
func TakeWhile(done <-chan struct{}, in <-chan interface{}, value interface{}) <-chan interface{} {
	return TakeWhileFn(done, in, func() interface{} { return value })
}

// TakeWhileFn -- 入力用チャネルから取得した値が指定した関数の戻り値と同一である間、値を取得し続けるチャネルを返します。
func TakeWhileFn(done <-chan struct{}, in <-chan interface{}, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func(fn func() interface{}) {
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
