package chans_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func TestLoopContext(t *testing.T) {
	// Arrange
	var (
		ctx   = context.Background()
		start = 1
		end   = 5
		out   = []int{1, 2, 3, 4, 5}
	)

	// Act
	var ret <-chan int = chans.LoopContext(ctx, start, end)

	// Assert
	for v := range chans.EnumerateContext(ctx, ret) {
		if out[v.Index] != v.Value {
			t.Errorf("[want] %v\t[got] %v", out[v.Index], v.Value)
		}
	}
}

func TestLoopInfiniteContext(t *testing.T) {
	// Arrange
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 100*time.Millisecond)
		interval = 10 * time.Millisecond
	)
	defer cxl()

	// Act
	var ret <-chan int = chans.LoopInfiniteContext(ctx)
	var out <-chan int = chans.IntervalContext(ctx, ret, interval)

	// Assert
	tmp := make([]int, 0)
	for v := range out {
		tmp = append(tmp, v)
	}

	if len(tmp) < 5 {
		t.Errorf("wrong count [%d]", len(tmp))
	}
}

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

			r := make([]int, 0)
			for v := range chans.Loop(done, c.in.start, c.in.end) {
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

			r := make([]int, 0)
			for v := range chans.LoopInfinite(mainCtx.Done()) {
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
