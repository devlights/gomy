package chans_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleForEach() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	for v := range chans.ForEach(procCtx.Done(), 1, 2, 3) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
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
