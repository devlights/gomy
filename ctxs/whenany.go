package ctxs

import (
	"context"

	"github.com/devlights/gomy/chans"
)

// WhenAny は、chans.WhenAny() の context.Context 版です.
//
// 指定したコンテキストのどれかが完了したら完了となる context.Context を生成して返します.
//
// Example:
//
//	<-ctxs.WhenAny(procCtx, ctx1, ctx2, ctx3).Done()
func WhenAny(pCtx context.Context, c ...context.Context) context.Context {
	ctx, cancel := context.WithCancel(pCtx)
	go func(dones []<-chan struct{}) {
		defer cancel()

		select {
		case <-ctx.Done():
		case <-chans.WhenAny(dones...):
		}
	}(ToDoneCh(c...))
	return ctx
}
