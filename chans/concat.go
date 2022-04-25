package chans

// Concat -- 指定されたチャネルのシーケンスを順に消費していく単一のチャネルを返します。
func Concat[T any](done <-chan struct{}, chList ...<-chan T) <-chan T {
	if len(chList) == 0 {
		c := make(chan T)
		close(c)

		return c
	}

	chSeq := make(chan (<-chan T), len(chList))
	func() {
		defer close(chSeq)

		for _, c := range chList {
			chSeq <- c
		}
	}()

	return Bridge(done, chSeq)
}
