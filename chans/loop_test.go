package chans

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	type (
		testin struct {
			start, end int
		}
		testout struct {
			result []int
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				start: 0,
				end:   1,
			},
			out: testout{result: []int{0}},
		},
		{
			in: testin{
				start: 0,
				end:   5,
			},
			out: testout{result: []int{0, 1, 2, 3, 4}},
		},
	}

	for i, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			r := make([]int, 0, 0)
			for v := range Loop(done, c.in.start, c.in.end) {
				t.Logf("[test-%02d] %v", i, v)
				r = append(r, v)
			}

			if !reflect.DeepEqual(c.out.result, r) {
				t.Errorf("want: %v\tgot: %v", c.out.result, r)
			}
		}()
	}
}

func TestLoopInfinite(t *testing.T) {
	type (
		testin struct {
			timeLimit time.Duration
		}
		testout struct {
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{timeLimit: 100 * time.Millisecond},
			out: testout{},
		},
	}

	for i, c := range cases {
		func() {
			mainCtx, cancel := context.WithTimeout(context.Background(), c.in.timeLimit)
			defer cancel()

			r := make([]int, 0, 0)
			for v := range LoopInfinite(mainCtx.Done()) {
				t.Logf("[test-%02d] %v", i, v)
				r = append(r, v)

				<-time.After(100 * time.Millisecond)
			}

			if len(r) == 0 {
				t.Error("want: non-zero list\tgot zero list")
			}
		}()
	}
}
