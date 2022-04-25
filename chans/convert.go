package chans

// ToString -- 入力用チャネルから値を取得し、文字列に変換するチャネルを返します。
//
// 文字列に変換することが出来なかった場合は、引数 failedValue を出力用チャネルに送ります。
func ToString[T any](done <-chan struct{}, in <-chan T, failedValue string) <-chan string {
	out := make(chan string)

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

				s, ok := any(v).(string)
				if !ok {
					s = failedValue
				}

				select {
				case out <- s:
				case <-done:
				}
			}
		}
	}()

	return out
}

// ToInt -- 入力用チャネルから値を取得し、数値に変換するチャネルを返します。
//
// 数値に変換することが出来なかった場合は、引数 failedValue を出力用チャネルに送ります。
func ToInt[T any](done <-chan struct{}, in <-chan T, failedValue int) <-chan int {
	out := make(chan int)

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

				s, ok := any(v).(int)
				if !ok {
					s = failedValue
				}

				select {
				case out <- s:
				case <-done:
				}
			}
		}
	}()

	return out
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
