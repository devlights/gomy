package iter_test

import (
	"log"
	"os"
	"testing"

	"github.com/devlights/gomy/iter"
)

func ExampleZip() {
	var (
		appLog = log.New(os.Stdout, "", 0)
	)

	var (
		a = []interface{}{1, 2}
		b = []interface{}{1, 2, 3}
	)

	for i, v := range iter.Zip(a, b) {
		appLog.Printf("%d: %v:%v", i, v.Item1, v.Item2)
	}

	// Output:
	// 0: 1:1
	// 1: 2:2
}

func TestZip(t *testing.T) {
	var (
		zero  = []interface{}{}
		two   = []interface{}{1, 2}
		three = []interface{}{1, 2, 3}
		four  = []interface{}{1, 2, 3, 4}
	)

	tests := []struct {
		name   string
		a, b   []interface{}
		length int
	}{
		{"nil, nil", nil, nil, 0},
		{"2,3", two, three, 2},
		{"3,4", three, four, 3},
		{"4,4", four, four, 4},
		{"4,2", four, two, 2},
		{"0,0", zero, zero, 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			items := iter.Zip(test.a, test.b)
			if len(items) != test.length {
				t.Errorf("[want] %d\t[got] %d", test.length, len(items))
			}
		})
	}
}
