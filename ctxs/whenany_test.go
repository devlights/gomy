package ctxs_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/ctxs"
)

func ExampleWhenAny() {
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, 1*time.Second)
	)

	defer mainCancel()
	defer procCancel()

	ctx1 := ctxs.StartGoroutine(procCtx, 100*time.Millisecond)
	ctx2 := ctxs.StartGoroutine(procCtx, 200*time.Millisecond)
	ctx3 := ctxs.StartGoroutine(procCtx, 300*time.Millisecond)

	start := time.Now()
	<-ctxs.WhenAny(procCtx, ctx1, ctx2, ctx3).Done()
	elapsed := time.Since(start)

	fmt.Printf("elapsed: %vmsec\n", elapsed.Milliseconds())

	// Output:
	// elapsed: 100msec
}

func TestWhenAny(t *testing.T) {
	cases := []struct {
		name   string
		delays []time.Duration
		limit  time.Duration
	}{
		{"100,200,300", []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}, 110 * time.Millisecond},
		{"100,500", []time.Duration{100 * time.Millisecond, 500 * time.Millisecond}, 110 * time.Millisecond},
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
				tasks = append(tasks, ctxs.StartGoroutine(procCtx, delay))
			}

			<-ctxs.WhenAny(procCtx, tasks...).Done()

			elapsed := time.Since(start)
			if c.limit < elapsed {
				t.Errorf("[want] less than %v\t[got] %v", c.limit, elapsed)
			}

			//t.Logf("[elapsed] %v", elapsed)
		})
	}
}
