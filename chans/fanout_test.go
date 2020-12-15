package chans

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

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
			wa := WhenAll(FanOut(
				done,
				ForEach(done, c.in.input...),
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
			wg := FanOutWg(
				done,
				ForEach(done, c.in.input...),
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
