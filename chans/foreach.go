package chans

// ForEach -- 指定されたデータを出力するチャネルを生成します。
func ForEach[T any](done <-chan struct{}, in ...T) <-chan T {
	out := make(chan T)

	go func(data []T) {
		defer close(out)

		for _, v := range data {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}(in)

	return out
}
