package chans

// Filter -- 入力データチャネル in から取得したデータを predicate に渡して 真 となったデータを返すチャネルを生成します。
func Filter[T any](done <-chan struct{}, in <-chan T, predicate func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

	ChLoop:
		for {
			select {
			case <-done:
				break ChLoop
			case v, ok := <-in:
				if !ok {
					break ChLoop
				}

				if predicate(v) {
					select {
					case out <- v:
					case <-done:
					}
				}
			}
		}
	}()

	return out
}
