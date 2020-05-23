package chans

import (
	"context"
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	type (
		testin struct {
			values []interface{}
		}
		testout struct {
			results []interface{}
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "empty",
			in:   testin{values: make([]interface{}, 0, 0)},
			out:  testout{results: make([]interface{}, 0, 0)},
		},
		{
			name: "1 to 5",
			in:   testin{values: []interface{}{1, 2, 3, 4, 5}},
			out:  testout{results: []interface{}{1, 2, 3, 4, 5}},
		},
		{
			name: "hello",
			in:   testin{values: []interface{}{"h", "e", "l", "l", "o"}},
			out:  testout{results: []interface{}{"h", "e", "l", "l", "o"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			results := make([]interface{}, 0, 0)

			rootCtx := context.Background()
			ctx, cancel := context.WithCancel(rootCtx)
			defer cancel()

			outCh := Generator(ctx.Done(), c.in.values...)
			for v := range outCh {
				results = append(results, v)
			}

			if !reflect.DeepEqual(c.out.results, results) {
				t.Errorf("[want] %v\t[got] %v", c.out.results, results)
			}
		})
	}
}
