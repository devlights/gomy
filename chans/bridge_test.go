package chans_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func TestBridge(t *testing.T) {
	cases := []struct {
		name string
		in   [][]interface{}
		out  []interface{}
	}{
		{
			name: "single",
			in:   [][]interface{}{{1, 2, 3}},
			out:  []interface{}{1, 2, 3},
		},
		{
			name: "multi",
			in:   [][]interface{}{{1, 2, 3}, {4, 5, 6}},
			out:  []interface{}{1, 2, 3, 4, 5, 6},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				ctx, cxl = context.WithTimeout(context.Background(), 50*time.Millisecond)
			)
			defer cxl()

			chSeq := make(chan (<-chan interface{}))
			go func() {
				defer close(chSeq)
				for _, v := range c.in {
					chSeq <- chans.Generator(ctx.Done(), v...)
				}
			}()

			var results []interface{}
			for v := range chans.Bridge(ctx.Done(), chSeq) {
				results = append(results, v)
			}

			if !reflect.DeepEqual(c.out, results) {
				t.Errorf("[want] %v\t[got] %v", c.out, results)
			}
		})
	}
}
