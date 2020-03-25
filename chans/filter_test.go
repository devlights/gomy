package chans

import (
	"reflect"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	type (
		testin struct {
			input     []interface{}
			predicate func(interface{}) bool
		}
		testout struct {
			result []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				input: []interface{}{"hello", "world"},
				predicate: func(v interface{}) bool {
					if s, ok := v.(string); ok {
						if strings.HasPrefix(s, "w") {
							return true
						}

						return false
					}

					return false
				},
			},
			out: testout{result: []interface{}{"world"}},
		},
		{
			in: testin{
				input: []interface{}{1, 2, 3, 4, 5},
				predicate: func(v interface{}) bool {
					if i, ok := v.(int); ok {
						if i <= 3 {
							return true
						}

						return false
					}

					return false
				},
			},
			out: testout{result: []interface{}{1, 2, 3}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{})
			go func() {
				defer close(inCh)
				for _, v := range c.in.input {
					inCh <- v
				}
			}()

			results := make([]interface{}, 0, 0)
			for v := range Filter(done, inCh, c.in.predicate) {
				t.Logf("[test-%02d] %v", caseIndex, v)
				results = append(results, v)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
