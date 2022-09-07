package chans

import (
	"context"
	"sync"
)

// WhenAllContext は、 WhenAll の context.Context 版です.
func WhenAllContext(pCtx context.Context, channels ...<-chan struct{}) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
		done     = WhenAll(channels...)
	)

	go func() {
		defer cxl()
		select {
		case <-pCtx.Done():
		case <-done:
		}
	}()

	return ctx
}

// WhenAll -- 指定した１つ以上のチャネルの全てが閉じられたら、閉じるチャネルを返します。
//
// チャネルを一つも渡さずに呼び出すと、既に close 済みのチャネルを返します。
func WhenAll(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		nilCh := make(chan struct{})
		close(nilCh)

		return nilCh
	case 1:
		return channels[0]
	}

	allDone := make(chan struct{})
	go func() {
		defer close(allDone)

		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, v := range channels {
			go func(ch <-chan struct{}) {
				defer wg.Done()
				<-ch
			}(v)
		}

		wg.Wait()
	}()

	return allDone
}
