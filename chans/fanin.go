package chans

// FanIn -- 指定されたチャネルリストをファンインするチャネルを返します。
func FanIn(done <-chan struct{}, channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	chList := make([]<-chan struct{}, 0)
	for _, in := range channels {
		chList = append(chList, func() <-chan struct{} {
			terminated := make(chan struct{})

			go func(in <-chan interface{}) {
				defer close(terminated)

				for v := range in {
					select {
					case <-done:
						return
					case out <- v:
					}
				}
			}(in)

			return terminated
		}())
	}

	go func(chList []<-chan struct{}) {
		defer close(out)
		<-WhenAll(chList...)
	}(chList)

	return out
}
