package chans

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	type (
		testin struct {
			input []interface{}
			fn    MapFunc
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
				fn: func(v interface{}) interface{} {
					return strings.ToUpper(v.(string))
				},
			},
			out: testout{result: []interface{}{"HELLO", "WORLD"}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			results := make([]interface{}, 0, 0)
			for m := range Map(done, ForEach(done, c.in.input...), c.in.fn) {

				if v, ok := m.(*MapValue); ok {
					t.Logf("[test-%02d] [%v] ==> [%v]", caseIndex, v.Before, v.After)
					results = append(results, v.After)
				}
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}