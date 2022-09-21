package chans_test

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestMapContext(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3}
		in     = chans.GeneratorContext(ctx, values...)
		out    = []string{"1", "2", "3"}
	)

	// Act
	var ret <-chan *chans.MapValue[int, string] = chans.MapContext(ctx, in, func(v int) string {
		return strconv.Itoa(v)
	})

	// Assert
	for v := range chans.EnumerateContext(ctx, ret) {
		var (
			idx = v.Index
			val = v.Value
		)

		if values[idx] != val.Before {
			t.Errorf("[want] %v\t[got] %v", values[idx], val.Before)
		}

		if out[idx] != val.After {
			t.Errorf("[want] %v\t[got] %v", out[idx], val.After)
		}
	}
}

func TestMap(t *testing.T) {
	type (
		testin struct {
			input []string
			fn    chans.MapFunc[string, string]
		}
		testout struct {
			result []string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				input: []string{"hello", "world"},
				fn: func(v string) string {
					return strings.ToUpper(v)
				},
			},
			out: testout{result: []string{"HELLO", "WORLD"}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			results := make([]string, 0)
			for v := range chans.Map(done, chans.Generator(done, c.in.input...), c.in.fn) {
				t.Logf("[test-%02d] [%v] ==> [%v]", caseIndex, v.Before, v.After)
				results = append(results, v.After)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
