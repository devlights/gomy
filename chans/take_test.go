package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleTake() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.ForEach(procCtx.Done(), 1, 2, 3, 4, 5)
	takes := chans.Take(procCtx.Done(), numbers, 3)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleTakeWhile() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.ForEach(procCtx.Done(), 1, 1, 1, 4, 1)
	takes := chans.TakeWhile(procCtx.Done(), numbers, 1)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 1
	// 1
}

func TestTake(t *testing.T) {
	type (
		testin struct {
			total int
			count int
		}
		testout struct {
			count int
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "total10,count0",
			in: testin{
				total: 10,
				count: 0,
			},
			out: testout{count: 0},
		},
		{
			name: "total10,count1",
			in: testin{
				total: 10,
				count: 1,
			},
			out: testout{count: 1},
		},
		{
			name: "total10,count5",
			in: testin{
				total: 10,
				count: 5,
			},
			out: testout{count: 5},
		},
		{
			name: "total10,count10",
			in: testin{
				total: 10,
				count: 10,
			},
			out: testout{count: 10},
		},
		{
			name: "total10,count12",
			in: testin{
				total: 10,
				count: 12,
			},
			out: testout{count: 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, c.in.total)
			func() {
				defer close(inCh)

				for i := 0; i < c.in.total; i++ {
					inCh <- i
				}
			}()

			takeCh := chans.Take(done, inCh, c.in.count)

			recvCount := 0
			for range takeCh {
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		})
	}
}

func TestTakeWhile(t *testing.T) {
	type (
		testin struct {
			value int
			data  []int
		}
		testout struct {
			count int
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "1-only",
			in: testin{
				data:  []int{1},
				value: 1,
			},
			out: testout{count: 1},
		},
		{
			name: "1, 1, 1, 1, 1, 2, 2",
			in: testin{
				data:  []int{1, 1, 1, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 5},
		},
		{
			name: "1, 1, 2, 1, 1, 2, 2",
			in: testin{
				data:  []int{1, 1, 2, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, len(c.in.data))
			func() {
				defer close(inCh)

				for _, v := range c.in.data {
					inCh <- v
				}
			}()

			takeCh := chans.TakeWhile(done, inCh, c.in.value)

			recvCount := 0
			for range takeCh {
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		})
	}
}

func TestTakeWhileFn(t *testing.T) {
	type (
		testin struct {
			fn   func() interface{}
			data []int
		}
		testout struct {
			count int
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "1-only",
			in: testin{
				data: []int{1},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 1},
		},
		{
			name: "1, 1, 1, 1, 1, 2, 2",
			in: testin{
				data: []int{1, 1, 1, 1, 1, 2, 2},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 5},
		},
		{
			name: "1, 1, 2, 1, 1, 2, 2",
			in: testin{
				data: []int{1, 1, 2, 1, 1, 2, 2},
				fn:   func() interface{} { return 1 },
			},
			out: testout{count: 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, len(c.in.data))
			func() {
				defer close(inCh)

				for _, v := range c.in.data {
					inCh <- v
				}
			}()

			takeCh := chans.TakeWhileFn(done, inCh, c.in.fn)

			recvCount := 0
			for range takeCh {
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		})
	}
}
