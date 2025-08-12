package chans_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"slices"

	"github.com/devlights/gomy/chans"
)

func TestFilterContext(t *testing.T) {
	// Arrange
	var (
		ctx       = context.Background()
		values    = []int{1, 2, 3, 4, 5}
		in        = chans.Generator(ctx.Done(), values...)
		out       = []int{4, 5}
		predicate = func(v int) bool {
			return v > 3
		}
	)

	// Act
	var ret <-chan int = chans.FilterContext(ctx, in, predicate)

	// Assert
	for v := range ret {
		if !slices.Contains(out, v) {
			t.Errorf("%v is not included in the %v", v, out)
		}
	}
}

func TestFilter(t *testing.T) {
	type (
		testin struct {
			input     []interface{}
			predicate func(interface{}) bool
		}
		testout struct {
			result []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				input: []interface{}{"hello", "world"},
				predicate: func(v interface{}) bool {
					if s, ok := v.(string); ok {
						return strings.HasPrefix(s, "w")
					}

					return false
				},
			},
			out: testout{result: []interface{}{"world"}},
		},
		{
			in: testin{
				input: []interface{}{1, 2, 3, 4, 5},
				predicate: func(v interface{}) bool {
					if i, ok := v.(int); ok {
						return i <= 3
					}

					return false
				},
			},
			out: testout{result: []interface{}{1, 2, 3}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{})
			go func() {
				defer close(inCh)
				for _, v := range c.in.input {
					inCh <- v
				}
			}()

			results := make([]interface{}, 0)
			for v := range chans.Filter(done, inCh, c.in.predicate) {
				t.Logf("[test-%02d] %v", caseIndex, v)
				results = append(results, v)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
