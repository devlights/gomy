package chans

import (
	"testing"
)

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
			out: testout{count: 0},
		},
		{
			in: testin{
				total: 10,
				count: 1,
			},
			out: testout{count: 1},
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
			out: testout{count: 10},
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

			takeCh := Take(done, inCh, c.in.count)

			recvCount := 0
			for v := range takeCh {
				t.Logf("[test-%02d][take] %v\n", caseCount+1, v)
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		}()
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
			out: testout{count: 1},
		},
		{
			in: testin{
				data:  []int{1, 1, 1, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 5},
		},
		{
			in: testin{
				data:  []int{1, 1, 2, 1, 1, 2, 2},
				value: 1,
			},
			out: testout{count: 2},
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

			takeCh := TakeWhile(done, inCh, c.in.value)

			recvCount := 0
			for v := range takeCh {
				t.Logf("[test-%02d][take] %v\n", caseCount+1, v)
				recvCount++
			}

			if c.out.count != recvCount {
				t.Errorf("want: %v\tgot: %v", c.out.count, recvCount)
			}
		}()
	}
}
