package chans_test

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
	"golang.org/x/exp/slices"
)

func TestFanOutContext(t *testing.T) {
	// Arrange
	var (
		ctx         = context.Background()
		values      = []int{1, 2, 3, 4, 5}
		in          = chans.Generator(ctx.Done(), values...)
		out         = make([]int, 0, len(values))
		workerCount = len(values)
	)

	// Act
	var ret []context.Context = chans.FanOutContext(ctx, in, workerCount, func(v int) {
		out = append(out, v)
		t.Log(out)
	})

	// Assert
	if len(ret) != workerCount {
		t.Errorf("[want] %v\t[got] %v", workerCount, len(ret))
	}

	<-ctxs.WhenAll(ctx, ret...).Done()

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	if !slices.Equal(values, out) {
		t.Errorf("[want] equal\t[got] %v, %v", values, out)
	}
}

func TestFanOutWgContext(t *testing.T) {
	// Arrange
	var (
		ctx         = context.Background()
		values      = []int{1, 2, 3, 4, 5}
		in          = chans.Generator(ctx.Done(), values...)
		out         = make([]int, 0, len(values))
		workerCount = len(values)
	)

	// Act
	var ret *sync.WaitGroup = chans.FanOutWgContext(ctx, in, workerCount, func(v int) {
		out = append(out, v)
		t.Log(out)
	})

	// Assert
	ret.Wait()

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	if !slices.Equal(values, out) {
		t.Errorf("[want] equal\t[got] %v, %v", values, out)
	}
}

func TestFanOut(t *testing.T) {
	type (
		testin struct {
			workerCount int
			input       []interface{}
			interval    time.Duration
		}
		testout struct {
			results    []int
			estimation time.Duration
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				workerCount: 1,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: (((6/1 + 1) * 10) + 20) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 2,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: (((6/2 + 1) * 10) + 20) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 3,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: (((6/3 + 1) * 10) + 20) * time.Millisecond,
			},
		},
	}

	for caseIndex, c := range cases {
		func(index int) {
			done := make(chan struct{})
			defer close(done)

			results := make(chan int, len(c.in.input))

			start := time.Now()
			wa := chans.WhenAll(chans.FanOut(
				done,
				chans.ForEach(done, c.in.input...),
				c.in.workerCount,
				func(v interface{}) {
					<-time.After(c.in.interval)

					if i, ok := v.(int); ok {
						results <- (i * 2)
					}
				})...)

			<-wa
			close(results)
			t.Logf("[workerCount=%d][estimation] %v\t[elapsed] %v", c.in.workerCount, c.out.estimation, time.Since(start))

			var values []int
			for v := range results {
				values = append(values, v)
			}

			sort.Ints(values)
			if !reflect.DeepEqual(c.out.results, values) {
				t.Errorf("[want] %v\t[got] %v", c.out.results, values)
			}
		}(caseIndex + 1)
	}
}

func TestFanOutWg(t *testing.T) {
	type (
		testin struct {
			workerCount int
			input       []interface{}
			interval    time.Duration
		}
		testout struct {
			results    []int
			estimation time.Duration
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				workerCount: 1,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: ((6/1 + 1) * 10) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 2,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: ((6/2 + 1) * 10) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 3,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 10 * time.Millisecond,
			},
			out: testout{
				results:    []int{2, 4, 6, 8, 10, 12},
				estimation: ((6/3 + 1) * 10) * time.Millisecond,
			},
		},
	}

	for caseIndex, c := range cases {
		func(index int) {
			done := make(chan struct{})
			defer close(done)

			results := make(chan int, len(c.in.input))

			start := time.Now()
			wg := chans.FanOutWg(
				done,
				chans.ForEach(done, c.in.input...),
				c.in.workerCount,
				func(v interface{}) {
					<-time.After(c.in.interval)

					if i, ok := v.(int); ok {
						results <- (i * 2)
					}
				})
			wg.Wait()
			close(results)

			t.Logf("[workerCount=%d][estimation] %v\t[elapsed] %v", c.in.workerCount, c.out.estimation, time.Since(start))

			var values []int
			for v := range results {
				values = append(values, v)
			}

			sort.Ints(values)
			if !reflect.DeepEqual(c.out.results, values) {
				t.Errorf("[want] %v\t[got] %v", c.out.results, values)
			}
		}(caseIndex + 1)
	}
}
