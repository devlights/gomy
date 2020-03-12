package mem

import (
	"testing"
)

func TestNewMem(t *testing.T) {
	type (
		testin struct {
			prefix string
		}
		testout struct {
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				prefix: "test",
			},
			out: testout{},
		},
	}

	for _, c := range cases {
		m := NewMem(
			Alloc(true),
			HeapAlloc(true),
			TotalAlloc(true),
			HeapObjects(true),
			Sys(true),
			NumGC(true))

		m.Print(c.in.prefix)
	}
}
