package chans

import (
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	type (
		testin struct {
			input    []interface{}
			interval time.Duration
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
				input:    []interface{}{1, 2, 3, 4, 5},
				interval: 100 * time.Millisecond,
			},
			out: testout{estimation: 600 * time.Millisecond},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			start := time.Now()
			for v := range Interval(done, ForEach(done, c.in.input...), c.in.interval) {
				t.Logf("[test-%02d] %v (%v)", caseIndex, v, time.Since(start))
			}

			end := time.Since(start)
			if c.out.estimation < end {
				t.Errorf("want: <= %v\tgot: %v", c.out.estimation, end)
			}
		}()
	}
}
