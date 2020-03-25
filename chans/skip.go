package chans

// Skip -- 指定した個数分、入力用チャネルから値をスキップするチャネルを返します。
func Skip(done <-chan struct{}, in <-chan interface{}, count int) <-chan interface{} {
	out := make(chan interface{})

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
func SkipWhile(done <-chan struct{}, in <-chan interface{}, value interface{}) <-chan interface{} {
	return SkipWhileFn(done, in, func() interface{} { return value })
}

// SkipWhileFn -- 入力用チャネルから取得した値が指定した関数の戻り値と同一である間、値をスキップし続けるチャネルを返します。
func SkipWhileFn(done <-chan struct{}, in <-chan interface{}, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func(fn func() interface{}) {
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
