package chans

// Bridge -- 指定されたチャネルのシーケンスを順に消費していく単一のチャネルを返します。
func Bridge(done <-chan struct{}, chanCh <-chan <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for {
			var ch <-chan interface{}
			select {
			case c, ok := <-chanCh:
				if !ok {
					return
				}

				ch = c
			case <-done:
				return
			}

			for v := range OrDone(done, ch) {
				select {
				case out <- v:
				case <-done:
				}
			}
		}
	}()

	return out
}
