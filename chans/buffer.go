package chans

// Buffer は、入力を指定した件数分に束ねてデータを返すチャネルを生成します.
func Buffer(done <-chan struct{}, in <-chan interface{}, count int) <-chan []interface{} {
	out := make(chan []interface{})

	go func() {
		defer close(out)

		for {
			items := make([]interface{}, 0, count)
			for item := range Take(done, in, count) {
				items = append(items, item)
			}

			if len(items) == 0 {
				break
			}

			select {
			case <-done:
			case out <- items:
			}

			if len(items) != count {
				break
			}
		}
	}()

	return out
}
