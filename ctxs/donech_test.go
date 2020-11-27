package ctxs_test

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/chans"
)

func ExampleToDoneCh() {
	const (
		goroutineCount = 3
	)

	var (
		iter = func(n int) []struct{}{return make([]struct{}, n)}
		g = func(pCtx context.Context, name string, timeout time.Duration) context.Context {
			ctx, cancel := context.WithTimeout(pCtx, timeout)
			go func() {
				defer cancel()
				fmt.Printf("[%s] start\n", name)
				time.Sleep(10 * time.Millisecond)
				fmt.Printf("[%s] done\n", name)
			}()
			return ctx
		}
	)

	var (
		rootCtx = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
	)

	defer mainCancel()

	contexts := make([]context.Context, goroutineCount)
	for i := range iter(goroutineCount) {
		contexts[i] = g(mainCtx, strconv.Itoa(i), 500 * time.Millisecond)
	}

	<-chans.WhenAll(ctxs.ToDoneCh(contexts...)...)

	fmt.Println("[main] done")

	// Unordered output:
	// [0] start
	// [1] start
	// [2] start
	// [2] done
	// [0] done
	// [1] done
	// [main] done
}