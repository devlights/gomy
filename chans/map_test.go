package chans_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestMap(t *testing.T) {
	type (
		testin struct {
			input []string
			fn    chans.MapFunc[string]
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
				input: []string{"hello", "world"},
				fn: func(v string) string {
					return strings.ToUpper(v)
				},
			},
			out: testout{result: []string{"HELLO", "WORLD"}},
		},
	}

	for caseIndex, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			results := make([]string, 0)
			for v := range chans.Map(done, chans.ForEach(done, c.in.input...), c.in.fn) {
				t.Logf("[test-%02d] [%v] ==> [%v]", caseIndex, v.Before, v.After)
				results = append(results, v.After)
			}

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
