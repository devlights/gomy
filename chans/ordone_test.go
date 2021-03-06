package chans_test

import (
	"context"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

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
