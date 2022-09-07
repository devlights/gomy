package chans

import "context"

// FanInContext は、 FanIn の context.Context 版です.
func FanInContext[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return FanIn(ctx.Done(), channels...)
}

// FanIn -- 指定されたチャネルリストをファンインするチャネルを返します。
func FanIn[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	out := make(chan T)

	chList := make([]<-chan struct{}, 0)
	for _, in := range channels {
		chList = append(chList, func() <-chan struct{} {
			terminated := make(chan struct{})

			go func(in <-chan T) {
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
