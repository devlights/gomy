package chans

import (
	"testing"
)

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

			repeatCh := Repeat(done, c.in.data...)

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

			repeatCh := RepeatFn(done, c.in.fn)

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
