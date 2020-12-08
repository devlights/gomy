package ctxs

import (
	"context"
	"time"
)

func StartGoroutine(pCtx context.Context, delay time.Duration) context.Context {
	ctx, cancel := context.WithCancel(pCtx)
	go func() {
		defer cancel()

		select {
		case <-ctx.Done():
		case <-time.After(delay):
		}
	}()
	return ctx
}
