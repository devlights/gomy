package chans_test

import (
	"context"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestTakeContext(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3, 4, 5}
		in     = chans.GeneratorContext(ctx, values...)
		out    = []int{1, 2, 3}
	)

	// Act
	var ret <-chan int = chans.TakeContext(ctx, in, 3)

	// Assert
	for v := range chans.EnumerateContext(ctx, ret) {
		var (
			idx = v.Index
			val = v.Value
		)

		if out[idx] != val {
			t.Errorf("[want] %v\t[got] %v", out[idx], val)
		}
	}
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

			inCh := make(chan int, len(c.in.data))
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
			fn   func() int
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
				fn:   func() int { return 1 },
			},
			out: testout{count: 1},
		},
		{
			name: "1, 1, 1, 1, 1, 2, 2",
			in: testin{
				data: []int{1, 1, 1, 1, 1, 2, 2},
				fn:   func() int { return 1 },
			},
			out: testout{count: 5},
		},
		{
			name: "1, 1, 2, 1, 1, 2, 2",
			in: testin{
				data: []int{1, 1, 2, 1, 1, 2, 2},
				fn:   func() int { return 1 },
			},
			out: testout{count: 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan int, len(c.in.data))
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
