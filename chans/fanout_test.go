package chans

import (
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
	type (
		testin struct {
			numGoroutine int
			input        []interface{}
			interval     time.Duration
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
				numGoroutine: 1,
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
				numGoroutine: 2,
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
				numGoroutine: 3,
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
				c.in.numGoroutine,
				func(v interface{}) {
					<-time.After(c.in.interval)
					// t.Logf("[test-%02d] %v", index, v)
				})...)
			elapsed := time.Since(start)

			t.Logf("[estimation] %v\t[elapsed] %v", c.out.estimation, elapsed)
			if c.out.estimation < elapsed {
				t.Errorf("want: <= %v\tgot: %v", c.out.estimation, elapsed)
			}
		}(caseIndex + 1)
	}
}
