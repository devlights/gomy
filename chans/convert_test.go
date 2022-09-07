package chans_test

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestConvertContext(t *testing.T) {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	var (
		in  = chans.Generator(ctx.Done(), 1, 2, 3, 4, 5)
		out = []string{"1", "2", "3", "4", "5"}
		fn  = func(i int) string { return strconv.Itoa(i) }
	)

	results := make([]string, 0)
	for v := range chans.ConvertContext(ctx, in, fn) {
		results = append(results, v)
	}

	if !reflect.DeepEqual(out, results) {
		t.Errorf("[want] %v\t[got] %v\n", out, results)
	}
}

func TestConvert(t *testing.T) {
	var (
		done = make(chan struct{})
		in   = chans.Generator(done, 1, 2, 3, 4, 5)
		out  = []string{"1", "2", "3", "4", "5"}
		fn   = func(i int) string { return strconv.Itoa(i) }
	)
	defer close(done)

	results := make([]string, 0)
	for v := range chans.Convert(done, in, fn) {
		results = append(results, v)
	}

	if !reflect.DeepEqual(out, results) {
		t.Errorf("[want] %v\t[got] %v\n", out, results)
	}
}

func TestToString(t *testing.T) {
	type (
		testin struct {
			data []interface{}
		}
		testout struct {
			result []string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data: []interface{}{"hello", "world"},
			},
			out: testout{
				result: []string{"hello", "world"},
			},
		},
		{
			in: testin{
				data: []interface{}{"hello", 100, "world"},
			},
			out: testout{
				result: []string{"hello", "", "world"},
			},
		},
	}

	for _, c := range cases {
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

			toStrCh := chans.ToString(done, inCh, "")

			results := make([]string, 0)
			for s := range toStrCh {
				results = append(results, s)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}

func TestToInt(t *testing.T) {
	type (
		testin struct {
			data []interface{}
		}
		testout struct {
			result []int
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data: []interface{}{"hello", "world"},
			},
			out: testout{
				result: []int{-1, -1},
			},
		},
		{
			in: testin{
				data: []interface{}{"hello", 100, "world"},
			},
			out: testout{
				result: []int{-1, 100, -1},
			},
		},
	}

	for _, c := range cases {
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

			toIntCh := chans.ToInt(done, inCh, -1)

			results := make([]int, 0)
			for i := range toIntCh {
				results = append(results, i)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}

func TestFromIntCh(t *testing.T) {
	inCh := make(chan int, 2)
	inCh <- 1
	inCh <- 2

	close(inCh)

	resultCh := chans.FromIntCh(inCh)

	v := <-resultCh
	t.Logf("[result] %T (%v)", v, v)

	v = <-resultCh
	t.Logf("[result] %T (%v)", v, v)
}

func TestFromStringCh(t *testing.T) {
	inCh := make(chan string, 2)
	inCh <- "hello"
	inCh <- "world"

	close(inCh)

	resultCh := chans.FromStringCh(inCh)

	v := <-resultCh
	t.Logf("[result] %T (%v)", v, v)

	v = <-resultCh
	t.Logf("[result] %T (%v)", v, v)
}
