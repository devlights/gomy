package ctxs

import (
	"context"

	"github.com/devlights/gomy/chans"
)

// WhenAll は、chans.WhenAll() の context.Context 版です.
//
// 以下の条件のいずれかが満たされた場合に完了となる context.Context を生成して返します.
//
//   - pCtx が完了
//   - pCtx 以外のコンテキストが全て完了
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
