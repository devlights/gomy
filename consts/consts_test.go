package consts

import (
	"testing"
)

func TestReturnCode(t *testing.T) {
	type (
		testin struct {
			value int
		}
		testout struct {
			value int
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "exitsuccess",
			in:   testin{value: ExitSuccess},
			out:  testout{value: 0},
		},
		{
			name: "exitfailure",
			in:   testin{value: ExitFailure},
			out:  testout{value: 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.in.value != c.out.value {
				t.Errorf("[want] %v\t[got] %v", c.out.value, c.in.value)
			}
		})
	}
}
