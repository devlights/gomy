package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleSkip() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 1, 1, 4, 5)
	items := chans.Skip(procCtx.Done(), numbers, 3)

	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// 4
	// 5
}

func TestSkip(t *testing.T) {
	type (
		testin struct {
			total int
			count int
		}
		testout struct {
			count int
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				total: 10,
				count: 0,
			},
			out: testout{count: 10},
		},
		{
			in: testin{
				total: 10,
				count: 1,
			},
			out: testout{count: 9},
		},
		{
			in: testin{
				total: 10,
				count: 5,
			},
			out: testout{count: 5},
		},
		{
			in: testin{
				total: 10,
				count: 10,
			},
			out: testout{count: 0},
		},
	}

	for caseCount, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, c.in.total)
			func() {
				defer close(inCh)

				for i := 0; i < c.in.total; i++ {
					inCh <- i
				}
			}()

			skipCh := chans.Skip(done, inCh, c.in.count)

			recvCount := 0
			for v := range skipCh {
				t.Logf("[test-%02d][skip] %v\n", caseCount+1, v)
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		}()
	}
}

func TestSkipWhile(t *testing.T) {
	type (
		testin struct {
			value int
			data  []int
		}
		testout struct {
			count int
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data:  []int{1},
				value: 1,
			},
			out: testout{count: 0},
		},
		{
			in: testin{
				data:  []int{1, 1, 1, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 2},
		},
		{
			in: testin{
				data:  []int{1, 1, 2, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 5},
		},
	}

	for caseCount, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, len(c.in.data))
			func() {
				defer close(inCh)

				for _, v := range c.in.data {
					inCh <- v
				}
			}()

			skipCh := chans.SkipWhile(done, inCh, c.in.value)

			recvCount := 0
			for v := range skipCh {
				t.Logf("[test-%02d][skip] %v\n", caseCount+1, v)
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		}()
	}
}

func TestSkipWhileFn(t *testing.T) {
	type (
		testin struct {
			fn   func() interface{}
			data []int
		}
		testout struct {
			count int
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data: []int{1},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 0},
		},
		{
			in: testin{
				data: []int{1, 1, 1, 1, 1, 2, 2},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 2},
		},
		{
			in: testin{
				data: []int{1, 1, 2, 1, 1, 2, 2},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 5},
		},
	}

	for caseCount, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, len(c.in.data))
			func() {
				defer close(inCh)

				for _, v := range c.in.data {
					inCh <- v
				}
			}()

			skipCh := chans.SkipWhileFn(done, inCh, c.in.fn)

			recvCount := 0
			for v := range skipCh {
				t.Logf("[test-%02d][skip] %v\n", caseCount+1, v)
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		}()
	}
}
