package chans

// Repeat -- 指定した値を永遠と繰り返すチャネルを返します。
func Repeat(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	out := make(chan interface{})

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
func RepeatFn(done <-chan struct{}, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})

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
