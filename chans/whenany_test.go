package chans_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleWhenAny() {
	fn := func(tlimit time.Duration) <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)
			time.Sleep(tlimit)
		}()

		return done
	}

	done1 := fn(100 * time.Millisecond)
	done2 := fn(200 * time.Millisecond)
	done3 := fn(300 * time.Millisecond)

	start := time.Now()
	<-chans.WhenAny(done1, done2, done3)
	elapsed := time.Since(start)

	fmt.Printf("elapsed: about 100msec ==> %v\n", elapsed <= 110*time.Millisecond)

	// Output:
	// elapsed: about 100msec ==> true
}

func TestWhenAny(t *testing.T) {
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
			out: testout{150 * time.Millisecond},
		},
		{
			in:  testin{makeChCount: 1},
			out: testout{150 * time.Millisecond},
		},
		{
			in:  testin{makeChCount: 2},
			out: testout{150 * time.Millisecond},
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
			if _, ok := <-chans.WhenAny(chList...); ok {
				t.Errorf("want: false\tgot: %v", ok)
			}

			elapsed := time.Since(start)
			t.Logf("len(ch)=%d\telapsed=%v\n", len(chList), elapsed)

			for _, v := range chList {
				ch := v
				<-ch
			}

			if c.out.limit < elapsed {
				t.Errorf("want: within %v\tgot %v", c.out.limit, elapsed)
			}
		}()
	}
}
