package iter_test

import (
	"testing"

	"github.com/devlights/gomy/iter"
)

func TestZip(t *testing.T) {
	var (
		two = []interface{}{1,2}
		three = []interface{}{1,2,3}
		four = []interface{}{1,2,3,4}
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
