package chans

// Concat -- 指定されたチャネルのシーケンスを順に消費していく単一のチャネルを返します。
func Concat(done <-chan struct{}, chList ...<-chan interface{}) <-chan interface{} {
	if len(chList) == 0 {
		c := make(chan interface{})
		close(c)

		return c
	}

	chSeq := make(chan (<-chan interface{}), len(chList))
	func() {
		defer close(chSeq)

		for _, c := range chList {
			chSeq <- c
		}
	}()

	return Bridge(done, chSeq)
}
