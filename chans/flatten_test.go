package chans

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	type (
		testin struct {
			data [][]interface{}
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
			in: testin{data: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
			}},
			out: testout{result: []interface{}{
				1, 2, 3, 4, 5, 6,
			}},
		},
	}

	for _, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			chList := make([]<-chan interface{}, 0, len(c.in.data))
			for _, list := range c.in.data {
				func() {
					ch := make(chan interface{}, len(list))
					defer close(ch)

					for _, v := range list {
						ch <- v
					}

					chList = append(chList, ch)
				}()
			}

			flattenCh := Flatten(done, chList...)

			results := make([]interface{}, 0, 0)
			for v := range flattenCh {
				t.Log(v)
				results = append(results, v)
			}

			// flatten の場合は、fanIn と異なり取得順序は確定なので中身も一致していることをテストする
			t.Logf("[c.out.result] %v", c.out.result)
			t.Logf("[results     ] %v", results)

			if !reflect.DeepEqual(c.out.result, results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
