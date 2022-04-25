package chans

// Convert -- 入力用チャネルから値を取得し変換するチャネルを返します。
func Convert[F any, T any](done <-chan struct{}, in <-chan F, fn func(F) T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case out <- fn(v):
				case <-done:
				}
			}
		}
	}()

	return out
}

// ToString -- 入力用チャネルから値を取得し、文字列に変換するチャネルを返します。
//
// 文字列に変換することが出来なかった場合は、引数 failedValue を出力用チャネルに送ります。
func ToString[T any](done <-chan struct{}, in <-chan T, failedValue string) <-chan string {
	return Convert(done, in, func(v T) string {
		s, ok := any(v).(string)
		if !ok {
			return failedValue
		}
		return s
	})
}

// ToInt -- 入力用チャネルから値を取得し、数値に変換するチャネルを返します。
//
// 数値に変換することが出来なかった場合は、引数 failedValue を出力用チャネルに送ります。
func ToInt[T any](done <-chan struct{}, in <-chan T, failedValue int) <-chan int {
	return Convert(done, in, func(v T) int {
		i, ok := any(v).(int)
		if !ok {
			return failedValue
		}
		return i
	})
}

// FromIntCh -- chan int を chan any に変換します。
func FromIntCh(ch <-chan int) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for v := range ch {
			out <- v
		}
	}()

	return out
}

// FromStringCh -- chan string を chan any に変換します。
func FromStringCh(ch <-chan string) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for v := range ch {
			out <- v
		}
	}()

	return out
}
