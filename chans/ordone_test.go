package chans_test

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/times"
)

func TestOrDoneContext(t *testing.T) {
	// Arrange
	var (
		rootCtx    = context.Background()
		ctx1, cxl1 = context.WithTimeout(rootCtx, 100*time.Millisecond)
		ctx2, cxl2 = context.WithTimeout(rootCtx, 300*time.Millisecond)
		in         = chans.LoopInfiniteContext(ctx2)
	)
	defer cxl1()
	defer cxl2()

	// Act
	elapsed := times.Stopwatch(func(start time.Time) {
		var ret <-chan int = chans.OrDoneContext(ctx1, in)
		for v := range ret {
			fmt.Fprint(io.Discard, v)
		}
	})

	<-ctxs.WhenAll(rootCtx, ctx1, ctx2).Done()

	// Assert
	if 300*time.Millisecond <= elapsed {
		t.Errorf("fail %v", elapsed)
	}
}

func TestOrDone(t *testing.T) {
	data := make([]interface{}, 0, 200)
	for i := 0; i < 200; i++ {
		data = append(data, i)
	}

	var (
		rootCtx     = context.Background()
		ctx, cancel = context.WithTimeout(rootCtx, 1*time.Millisecond)
		results     []interface{}
	)

	defer cancel()

	for v := range chans.OrDone(ctx.Done(), chans.Generator(ctx.Done(), data...)) {
		t.Logf("[result] %v", v)
		results = append(results, v)
	}

	if len(results) == 0 {
		t.Errorf("want: not 0\tgot: %v", len(results))
	}
}
