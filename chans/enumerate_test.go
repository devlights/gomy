package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleEnumerate() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 9, 8, 7)
	values := chans.Enumerate(procCtx.Done(), numbers)

	for e := range values {
		if v, ok := e.(*chans.IterValue); ok {
			fmt.Printf("%d:%v\n", v.Index, v.Value)
		}
	}

	// Output:
	// 0:9
	// 1:8
	// 2:7
}

func TestEnumerate(t *testing.T) {
	type (
		resultValue struct {
			index int
			value interface{}
		}
		testin struct {
			input []interface{}
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
			in: testin{input: []interface{}{"hello", "world"}},
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

			inCh := make(chan interface{})
			go func() {
				defer close(inCh)
				for _, v := range c.in.input {
					inCh <- v
				}
			}()

			for e := range chans.Enumerate(done, inCh) {
				if v, ok := e.(*chans.IterValue); ok {
					t.Logf("[test-%02d] [%v][%v]", caseIndex, v.Index, v.Value)

					r := c.out.output[v.Index]
					if r.index != v.Index {
						t.Errorf("want: index %v\tgot: index %v", r.index, v.Index)
					}

					if r.value != v.Value {
						t.Errorf("want value %v\tgot: value %v", r.value, v.Value)
					}
				}
			}
		}()
	}
}
