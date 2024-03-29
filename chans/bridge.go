package chans

import "context"

// BridgeContext-- Bridge の context.Context 版です。
func BridgeContext[T any](ctx context.Context, chanCh <-chan <-chan T) <-chan T {
	return Bridge(ctx.Done(), chanCh)
}

// Bridge -- 指定されたチャネルのシーケンスを順に消費していく単一のチャネルを返します。
func Bridge[T any](done <-chan struct{}, chanCh <-chan <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			var ch <-chan T
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
