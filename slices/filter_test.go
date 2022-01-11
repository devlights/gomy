package slices_test

import (
	"strings"
	"testing"

	"github.com/devlights/gomy/slices"
)

func TestFilter(t *testing.T) {
	var (
		odd = func(v interface{}) bool {
			return v.(int)%2 == 0
		}
		str = func(v interface{}) bool {
			return strings.HasPrefix(v.(string), "w")
		}
	)

	tests := []struct {
		name            string
		in, out, postIn []interface{}
		fn              func(v interface{}) bool
	}{
		{
			"odd",
			[]interface{}{1, 2, 3}, []interface{}{2}, []interface{}{1, 2, 3},
			odd,
		},
		{
			"str",
			[]interface{}{"hello", "world"}, []interface{}{"world"}, []interface{}{"hello", "world"},
			str,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := slices.Filter(test.in, test.fn)
			for i := 0; i < len(test.out); i++ {
				if test.out[i] != got[i] {
					t.Errorf("[test.out][want] %v\t[got] %v", test.out[i], got[i])
				}
			}
			for i := 0; i < len(test.in); i++ {
				if test.in[i] != test.postIn[i] {
					t.Errorf("[test.postIn][want] %v\t[got] %v", test.in[i], test.postIn[i])
				}
			}
		})
	}
}

func TestFilterD(t *testing.T) {
	var (
		odd = func(v interface{}) bool {
			return v.(int)%2 == 0
		}
		str = func(v interface{}) bool {
			return strings.HasPrefix(v.(string), "w")
		}
	)

	tests := []struct {
		name            string
		in, out, postIn []interface{}
		fn              func(v interface{}) bool
	}{
		{
			"odd",
			[]interface{}{1, 2, 3}, []interface{}{2}, []interface{}{2, 2, 3},
			odd,
		},
		{
			"str",
			[]interface{}{"hello", "world"}, []interface{}{"world"}, []interface{}{"world", "world"},
			str,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := slices.FilterD(test.in, test.fn)
			for i := 0; i < len(test.out); i++ {
				if test.out[i] != got[i] {
					t.Errorf("[test.out][want] %v\t[got] %v", test.out[i], got[i])
				}
			}
			for i := 0; i < len(test.in); i++ {
				if test.in[i] != test.postIn[i] {
					t.Errorf("[test.postIn][want] %v\t[got] %v", test.in[i], test.postIn[i])
				}
			}
		})
	}
}
