package chans

// ToString -- 入力用チャネルから値を取得し、文字列に変換するチャネルを返します。
//
// 文字列に変換することが出来なかった場合は、引数 failedValue を出力用チャネルに送ります。
func ToString(done <-chan struct{}, in <-chan interface{}, failedValue string) <-chan string {
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

				s, ok := v.(string)
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
func ToInt(done <-chan struct{}, in <-chan interface{}, failedValue int) <-chan int {
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

				s, ok := v.(int)
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

// FromIntCh -- chan int を chan interface{} に変換します。
func FromIntCh(ch <-chan int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for v := range ch {
			out <- v
		}
	}()

	return out
}

// FromStringCh -- chan string を chan interface{} に変換します。
func FromStringCh(ch <-chan string) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for v := range ch {
			out <- v
		}
	}()

	return out
}
