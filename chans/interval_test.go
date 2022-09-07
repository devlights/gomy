package chans_test

import (
	"context"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/times"
)

func TestIntervalContext(t *testing.T) {
	// Arrange
	var (
		ctx      = context.Background()
		values   = []int{1, 2, 3}
		in       = chans.GeneratorContext(ctx, values...)
		interval = 100 * time.Millisecond
	)

	// Act
	var ret <-chan int = chans.IntervalContext(ctx, in, interval)

	elapsed := times.Stopwatch(func(start time.Time) {
		for v := range ret {
			t.Log(v)
		}
	})

	// Assert
	if !(300*time.Millisecond <= elapsed && elapsed <= 1*time.Second) {
		t.Errorf("invalid interval flow [%v]", elapsed)
	}
}

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
				interval: 10 * time.Millisecond,
			},
			out: testout{estimation: (10*5 + 10) * time.Millisecond},
		},
		{
			in: testin{
				input:    []interface{}{1, 2, 3, 4, 5},
				interval: 100 * time.Millisecond,
			},
			out: testout{estimation: (100*5 + 10) * time.Millisecond},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			start := time.Now()
			var lastElapsed time.Duration
			for v := range chans.Interval(done, chans.ForEach(done, c.in.input...), c.in.interval) {
				lastElapsed = time.Since(start)
				t.Logf("[test-%02d] %v (%v)", caseIndex, v, lastElapsed)
			}

			if c.out.estimation < lastElapsed {
				t.Errorf("want: <= %v\tgot: %v", c.out.estimation, lastElapsed)
			}
		}()
	}
}
