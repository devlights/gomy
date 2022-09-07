package chans_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestForEachContext(t *testing.T) {
	// Arrange
	var (
		ctx = context.Background()
		in  = []int{1, 2, 3, 4, 5}
	)

	// Act
	var ret <-chan int = chans.ForEachContext(ctx, in...)

	// Assert
	for v := range chans.EnumerateContext(ctx, ret) {
		if in[v.Index] != v.Value {
			t.Errorf("[want] %v\t[got] %v", in[v.Index], v.Index)
		}
	}
}

func TestForEach(t *testing.T) {
	type (
		testin struct {
			input []interface{}
		}
		testout struct {
			output []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{input: []interface{}{"hello", "world"}},
			out: testout{output: []interface{}{"hello", "world"}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			results := make([]interface{}, 0)
			for v := range chans.ForEach(done, c.in.input...) {
				t.Logf("[test-%02d] %v", caseIndex, v)
				results = append(results, v)
			}

			if !reflect.DeepEqual(c.out.output, results) {
				t.Errorf("want: %v\tgot: %v", c.out.output, results)
			}
		}()
	}
}
