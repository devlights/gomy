package deepcopy

import (
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
)

func TestGobCopy(t *testing.T) {
	type (
		testin struct {
			from interface{}
			to   interface{}
		}
		testout struct {
			valueEqualFn      func(x, y interface{}) bool
			pointerNotEqualFn func(x, y interface{}) bool
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "deepcopy(string)",
			in:   testin{from: "helloworld", to: ""},
			out: testout{
				valueEqualFn: func(x, y interface{}) bool {
					v1 := x.(string)
					v2 := y.(string)

					return v1 == v2
				},
				pointerNotEqualFn: func(x, y interface{}) bool {
					v1 := x.(string)
					v2 := y.(string)

					p1 := fmt.Sprintf("%p", &v1)
					p2 := fmt.Sprintf("%p", &v2)

					return p1 != p2
				},
			},
		},
		{
			name: "deepcopy(slice)",
			in:   testin{from: []int{1, 2, 3}, to: []int{}},
			out: testout{
				valueEqualFn: func(x, y interface{}) bool {
					v1 := x.([]int)
					v2 := y.([]int)

					return reflect.DeepEqual(v1, v2)
				},
				pointerNotEqualFn: func(x, y interface{}) bool {
					v1 := x.([]int)
					v2 := y.([]int)

					p1 := fmt.Sprintf("%p", &v1)
					p2 := fmt.Sprintf("%p", &v2)

					return p1 != p2
				},
			},
		},
		{
			name: "deepcopy(map)",
			in:   testin{from: map[int]string{1: "apple"}, to: make(map[int]string)},
			out: testout{
				valueEqualFn: func(x, y interface{}) bool {
					v1 := x.(map[int]string)
					v2 := y.(map[int]string)

					return reflect.DeepEqual(v1, v2)
				},
				pointerNotEqualFn: func(x, y interface{}) bool {
					v1 := x.(map[int]string)
					v2 := y.(map[int]string)

					p1 := fmt.Sprintf("%p", &v1)
					p2 := fmt.Sprintf("%p", &v2)

					return p1 != p2
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// https://qiita.com/juntaki/items/a75dfe7c8597557ed577
			// https://stackoverflow.com/questions/21934730/gob-type-not-registered-for-interface-mapstringinterface
			gob.Register(c.in.from)

			if err := GobCopy(&c.in.from, &c.in.to); err != nil {
				t.Errorf("%v", err)
			}

			if !c.out.valueEqualFn(c.in.from, c.in.to) {
				t.Errorf("[valueEqualFn][want] true\t[got] false")
			}
			if !c.out.pointerNotEqualFn(c.in.from, c.in.to) {
				t.Errorf("[pointerNotEqualFn][want] true\t[got] false")
			}
		})
	}
}
