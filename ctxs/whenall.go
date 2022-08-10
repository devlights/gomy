package ctxs

import (
	"context"

	"github.com/devlights/gomy/chans"
)

// WhenAll は、chans.WhenAll() の context.Context 版です.
//
// 指定したコンテキストが全て完了したら完了となる context.Context を生成して返します.
//
// Example:
//
//	<-ctxs.WhenAll(procCtx, ctx1, ctx2, ctx3).Done()
func WhenAll(pCtx context.Context, c ...context.Context) context.Context {
	ctx, cancel := context.WithCancel(pCtx)
	go func(dones []<-chan struct{}) {
		defer cancel()

		select {
		case <-ctx.Done():
		case <-chans.WhenAll(dones...):
		}
	}(ToDoneCh(c...))
	return ctx
}
