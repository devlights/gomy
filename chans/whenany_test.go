package chans

import (
	"testing"
	"time"
)

func TestWhenAny(t *testing.T) {
	type (
		testin struct {
			makeChCount int
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
			in:  testin{makeChCount: 0},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 1},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 2},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 3},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 4},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 5},
			out: testout{},
		},
		{
			in:  testin{makeChCount: 6},
			out: testout{},
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
			if _, ok := <-WhenAny(chList...); ok {
				t.Errorf("want: false\tgot: %v", ok)
			}

			t.Logf("len(ch)=%d\telapsed=%v\n", len(chList), time.Since(start))

			for _, v := range chList {
				<-v
			}
		}()
	}
}
