package chans

// ForEach -- 指定されたデータを出力するチャネルを生成します。
func ForEach(done <-chan struct{}, in ...interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func(data []interface{}) {
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
