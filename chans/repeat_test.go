package chans_test

import (
	"context"
	"testing"
	"time"

	"github.com/devlights/gomy/chans"
)

func TestRepeatContext(t *testing.T) {
	// Arrange
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 100*time.Millisecond)
		values   = []int{1, 2, 3}
	)
	defer cxl()

	// Act
	var ret <-chan int = chans.RepeatContext(ctx, values...)
	var out <-chan int = chans.IntervalContext(ctx, ret, 10*time.Millisecond)

	// Assert
	cnt := 0
	tmp := make(map[int]struct{})
	for v := range out {
		tmp[v] = struct{}{}
		cnt++
	}

	if cnt <= len(tmp) {
		t.Errorf("[want] more than %v\t[got] %v", len(tmp), cnt)
	}

	if len(tmp) != 3 {
		t.Errorf("[want] 3\t[got] %v", len(tmp))
	}
}

func TestRepeatFnContext(t *testing.T) {
	// Arrange
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 100*time.Millisecond)
		fn       = func() func() int {
			cnt := 1
			return func() int {
				v := cnt
				cnt++

				return v
			}
		}()
	)
	defer cxl()

	// Act
	var ret <-chan int = chans.RepeatFnContext(ctx, fn)
	var out <-chan int = chans.IntervalContext(ctx, ret, 10*time.Millisecond)

	// Assert
	tmp := make([]int, 0)
	for v := range out {
		tmp = append(tmp, v)
	}

	if len(tmp) <= 5 {
		t.Errorf("[want] 10\t[got] %v", len(tmp))
	}
}

func TestRepeat(t *testing.T) {
	type (
		testin struct {
			data []interface{}
		}
		testout struct {
			value []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data: []interface{}{1},
			},
			out: testout{
				value: []interface{}{1, 1, 1},
			},
		},
		{
			in: testin{
				data: []interface{}{1, 2, 3},
			},
			out: testout{
				value: []interface{}{1, 2, 3},
			},
		},
	}

	for index, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			repeatCh := chans.Repeat(done, c.in.data...)

			result1 := <-repeatCh
			result2 := <-repeatCh
			result3 := <-repeatCh

			if result1 != c.out.value[0] {
				t.Errorf("[testcase-%02d][result1] want: %v\tgot: %v", index+1, c.out.value[0], result1)
			}

			if result2 != c.out.value[1] {
				t.Errorf("[testcase-%02d][result2] want: %v\tgot: %v", index+1, c.out.value[1], result1)
			}

			if result3 != c.out.value[2] {
				t.Errorf("[testcase-%02d][result3] want: %v\tgot: %v", index+1, c.out.value[2], result1)
			}
		}()
	}
}

func TestRepeatFn(t *testing.T) {
	type (
		testin struct {
			fn func() interface{}
		}
		testout struct {
			value []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				fn: func() interface{} {
					return "helloworld"
				},
			},
			out: testout{
				value: []interface{}{"helloworld", "helloworld", "helloworld"},
			},
		},
	}

	for index, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			repeatCh := chans.RepeatFn(done, c.in.fn)

			result1 := <-repeatCh
			result2 := <-repeatCh
			result3 := <-repeatCh

			if result1 != c.out.value[0] {
				t.Errorf("[testcase-%02d][result1] want: %v\tgot: %v", index+1, c.out.value[0], result1)
			}

			if result2 != c.out.value[1] {
				t.Errorf("[testcase-%02d][result2] want: %v\tgot: %v", index+1, c.out.value[1], result1)
			}

			if result3 != c.out.value[2] {
				t.Errorf("[testcase-%02d][result3] want: %v\tgot: %v", index+1, c.out.value[2], result1)
			}
		}()
	}
}
