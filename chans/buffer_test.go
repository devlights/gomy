package chans_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/devlights/gomy/chans"
)

func ExampleBuffer() {
	var (
		done  = make(chan struct{})
		data  = []interface{}{1, 2, 3, 4, 5, 6, 7}
		count = 3
	)

	defer close(done)

	for chunk := range chans.Buffer(done, chans.Generator(done, data...), count) {
		fmt.Println(chunk)
	}

	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7]
}

func TestBuffer(t *testing.T) {
	cases := []struct {
		name  string
		in    []interface{}
		count int
		out   [][]interface{}
	}{
		{"src[1,2,3]count[1]", []interface{}{1, 2, 3}, 1, [][]interface{}{{1}, {2}, {3}}},
		{"src[1,2,3]count[3]", []interface{}{1, 2, 3}, 3, [][]interface{}{{1, 2, 3}}},
		{"src[1,2,3]count[2]", []interface{}{1, 2, 3}, 2, [][]interface{}{{1, 2}, {3}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			results := make([][]interface{}, 0)
			for chunk := range chans.Buffer(done, chans.ForEach(done, c.in...), c.count) {
				results = append(results, chunk)
			}

			for i, chunk := range results {
				if !reflect.DeepEqual(chunk, c.out[i]) {
					t.Errorf("[want] %v\t[got] %v -- [index] %d", c.out[i], chunk, i)
				}
			}
		})
	}
}
