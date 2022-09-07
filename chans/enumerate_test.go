package chans_test

import (
	"context"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestEnumerateContext(t *testing.T) {
	// Arrange
	var (
		values = []string{"hello", "world"}
		ctx    = context.Background()
		in     = chans.Generator(ctx.Done(), values...)
	)

	// Act
	var ret <-chan *chans.IterValue[string] = chans.EnumerateContext(ctx, in)

	// Assert
	for v := range ret {
		i := v.Index
		if v.Value != values[i] {
			t.Errorf("[want] %v\t[got] %v", values[i], v.Value)
		}
	}
}

func TestEnumerate(t *testing.T) {
	type (
		resultValue struct {
			index int
			value string
		}
		testin struct {
			input []string
		}
		testout struct {
			output []resultValue
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{input: []string{"hello", "world"}},
			out: testout{output: []resultValue{
				{
					index: 0,
					value: "hello",
				},
				{
					index: 1,
					value: "world",
				},
			}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan string)
			go func() {
				defer close(inCh)
				for _, v := range c.in.input {
					inCh <- v
				}
			}()

			for v := range chans.Enumerate(done, inCh) {
				t.Logf("[test-%02d] [%v][%v]", caseIndex, v.Index, v.Value)

				r := c.out.output[v.Index]
				if r.index != v.Index {
					t.Errorf("want: index %v\tgot: index %v", r.index, v.Index)
				}

				if r.value != v.Value {
					t.Errorf("want value %v\tgot: value %v", r.value, v.Value)
				}
			}
		}()
	}
}
