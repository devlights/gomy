package chans

// Tee -- Unix の tee コマンドのように一つの入力を２つに複製するチャネルを返します。
//
// noinspection GoNilness
func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range OrDone(done, in) {
			var ch1, ch2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case ch1 <- v:
					ch1 = nil
				case ch2 <- v:
					ch2 = nil
				}
			}
		}
	}()

	return out1, out2
}
