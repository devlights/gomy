package chans

import (
	"context"
	"sync"
)

func FanOutContext[T any](ctx context.Context, in <-chan T, workerCount int, callback func(T)) []context.Context {
	var (
		dones   = FanOut(ctx.Done(), in, workerCount, callback)
		results = make([]context.Context, len(dones))
	)

	for i, done := range dones {
		results[i] = func(done <-chan struct{}) context.Context {
			var (
				ctx, cxl = context.WithCancel(ctx)
			)

			go func() {
				defer cxl()
				<-done
			}()

			return ctx
		}(done)
	}

	return results
}

// FanOut -- 指定されたチャネルの処理を指定されたワーカーの数でファンアウトします。
//
// チャネルからデータを取得するたびに引数 callback が呼ばれます。
func FanOut[T any](done <-chan struct{}, in <-chan T, workerCount int, callback func(T)) []<-chan struct{} {
	outChList := make([]<-chan struct{}, 0)

	for i := 0; i < workerCount; i++ {
		out := make(chan struct{})

		go func() {
			defer close(out)

			for v := range OrDone(done, in) {
				select {
				case <-done:
					return
				default:
				}

				callback(v)
			}
		}()

		outChList = append(outChList, out)
	}

	return outChList
}

// FanOutWgContext は、FanOutWg の context.Context 版です.
func FanOutWgContext[T any](ctx context.Context, in <-chan T, workerCount int, callback func(T)) *sync.WaitGroup {
	return FanOutWg(ctx.Done(), in, workerCount, callback)
}

// FanOutWg -- FanOut() の sync.WaitGroup を返す版です。
func FanOutWg[T any](done <-chan struct{}, in <-chan T, workerCount int, callback func(T)) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range OrDone(done, in) {
				select {
				case <-done:
					return
				default:
				}

				callback(v)
			}
		}()
	}

	return &wg
}
