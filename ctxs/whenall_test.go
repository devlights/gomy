package ctxs_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/ctxs"
)

func ExampleWhenAll() {
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, 1*time.Second)
	)

	defer mainCancel()
	defer procCancel()

	ctx1 := g(procCtx, 100*time.Millisecond)
	ctx2 := g(procCtx, 200*time.Microsecond)
	ctx3 := g(procCtx, 300*time.Millisecond)

	start := time.Now()
	<-ctxs.WhenAll(procCtx, ctx1, ctx2, ctx3).Done()
	elapsed := time.Since(start)

	fmt.Printf("elapsed: %vmsec\n", elapsed.Milliseconds())

	// Output:
	// elapsed: 300msec
}

func TestWhenAll(t *testing.T) {
	cases := []struct {
		name   string
		delays []time.Duration
		limit  time.Duration
	}{
		{"100,200,300", []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}, 400 * time.Millisecond},
		{"100,500", []time.Duration{100 * time.Millisecond, 500 * time.Millisecond}, 600 * time.Millisecond},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				rootCtx             = context.Background()
				mainCtx, mainCancel = context.WithCancel(rootCtx)
				procCtx, procCancel = context.WithTimeout(mainCtx, 1*time.Second)
			)

			defer mainCancel()
			defer procCancel()

			start := time.Now()

			var tasks []context.Context
			for _, delay := range c.delays {
				tasks = append(tasks, g(procCtx, delay))
			}

			<-ctxs.WhenAll(procCtx, tasks...).Done()

			elapsed := time.Since(start)
			if c.limit < elapsed {
				t.Errorf("[want] less than %v\t[got] %v", c.limit, elapsed)
			}

			//t.Logf("[elapsed] %v", elapsed)
		})
	}
}

func g(pCtx context.Context, delay time.Duration) context.Context {
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
