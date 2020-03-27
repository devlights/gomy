package chans

import (
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
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/1 + 1) * 100) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 2,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/2 + 1) * 100) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 3,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/3 + 1) * 100) * time.Millisecond,
			},
		},
	}

	for caseIndex, c := range cases {
		func(index int) {
			done := make(chan struct{})
			defer close(done)

			start := time.Now()
			<-WhenAll(FanOut(
				done,
				ForEach(done, c.in.input...),
				c.in.workerCount,
				func(v interface{}) {
					<-time.After(c.in.interval)
				})...)
			elapsed := time.Since(start)

			t.Logf("[workerCount=%d][estimation] %v\t[elapsed] %v", c.in.workerCount, c.out.estimation, elapsed)
			if c.out.estimation < elapsed {
				t.Errorf("want: <= %v\tgot: %v", c.out.estimation, elapsed)
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
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/1 + 1) * 100) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 2,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/2 + 1) * 100) * time.Millisecond,
			},
		},
		{
			in: testin{
				workerCount: 3,
				input: []interface{}{
					1, 2, 3, 4, 5, 6,
				},
				interval: 100 * time.Millisecond,
			},
			out: testout{
				estimation: ((6/3 + 1) * 100) * time.Millisecond,
			},
		},
	}

	for caseIndex, c := range cases {
		func(index int) {
			done := make(chan struct{})
			defer close(done)

			start := time.Now()
			wg := FanOutWg(
				done,
				ForEach(done, c.in.input...),
				c.in.workerCount,
				func(v interface{}) {
					<-time.After(c.in.interval)
				})
			wg.Wait()
			elapsed := time.Since(start)

			t.Logf("[workerCount=%d][estimation] %v\t[elapsed] %v", c.in.workerCount, c.out.estimation, elapsed)
			if c.out.estimation < elapsed {
				t.Errorf("want: <= %v\tgot: %v", c.out.estimation, elapsed)
			}
		}(caseIndex + 1)
	}
}
