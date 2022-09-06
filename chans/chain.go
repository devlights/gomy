package chans

import (
	"context"
	"time"
)

// ChainContext -- Chain の context.Context 版です.
func ChainContext(ctx context.Context, base context.Context, next func(finished time.Time)) context.Context {
	var (
		pCtx, pCxl = context.WithCancel(ctx)
		done       = Chain(ctx.Done(), base.Done(), next)
	)

	go func() {
		defer pCxl()
		<-done
	}()

	return pCtx
}

// Chain -- 指定された base チャネルがクローズした後に next で指定された関数を呼び出します.
//
// done チャネルがクローズした場合、 next は実行されません。
func Chain(done, base <-chan struct{}, next func(finished time.Time)) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		select {
		case <-done:
			return
		case <-base:
			if next != nil {
				next(time.Now())
			}

			return
		}
	}()

	return ch
}
