package chans

import (
	"testing"
	"time"
)

func TestWhenAll(t *testing.T) {
	type (
		testin struct {
			makeChCount int
		}
		testout struct {
			limit time.Duration
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{makeChCount: 0},
			out: testout{100 * time.Millisecond},
		},
		{
			in:  testin{makeChCount: 1},
			out: testout{150 * time.Millisecond},
		},
		{
			in:  testin{makeChCount: 2},
			out: testout{250 * time.Millisecond},
		},
	}

	makeCh := func(closeAfter time.Duration) <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			defer close(ch)
			time.Sleep(closeAfter)
		}()

		return ch
	}

	for _, c := range cases {
		func() {
			chList := make([]<-chan struct{}, 0, c.in.makeChCount)
			for i := 0; i < c.in.makeChCount; i++ {
				ch := makeCh(time.Duration((i+1)*100) * time.Millisecond)
				chList = append(chList, ch)
			}

			start := time.Now()
			if _, ok := <-WhenAll(chList...); ok {
				t.Errorf("want: false\tgot: %v", ok)
			}

			elapsed := time.Since(start)
			t.Logf("len(ch)=%d\telapsed=%v\n", len(chList), elapsed)

			if c.out.limit < elapsed {
				t.Errorf("want: within %v\tgot %v", c.out.limit, elapsed)
			}
		}()
	}
}
