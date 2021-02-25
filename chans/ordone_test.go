package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleOrDone() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Minute)
		genCtx, genCxl   = context.WithCancel(mainCtx)
	)

	defer mainCxl()
	defer procCxl()
	defer genCxl()

	inCh := chans.Generator(genCtx.Done(), "h", "e", "l", "l", "o")

	var result []interface{}
	for v := range chans.OrDone(procCtx.Done(), inCh) {
		func() {
			defer procCxl()
			result = append(result, v)
		}()
	}

	fmt.Printf("len(result) <= 2: %v", len(result) <= 2)

	// Output:
	// len(result) <= 2: true
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
