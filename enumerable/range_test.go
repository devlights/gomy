package enumerable

import (
	"testing"
)

type (
	testin struct {
		start, end int
	}

	testout struct {
		start, end, current int
	}

	testcase struct {
		in  testin
		out testout
	}
)

func TestEnumerableRange_InitialState(t *testing.T) {
	cases := []testcase{
		{
			in: testin{
				start: 0,
				end:   1,
			},
			out: testout{
				start:   0,
				end:     1,
				current: 0,
			},
		},
		{
			in: testin{
				start: 1,
				end:   2,
			},
			out: testout{
				start:   1,
				end:     2,
				current: 1,
			},
		},
	}

	for _, c := range cases {
		sut := NewRange(c.in.start, c.in.end)

		want := c.out.start
		got := sut.Start()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}

		want = c.out.end
		got = sut.End()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}

		want = c.out.current
		got = sut.Current()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}
	}
}

func TestEnumerableRange_Next(t *testing.T) {
	cases := []testcase{
		{
			in: testin{
				start: 0,
				end:   1,
			},
			out: testout{
				start:   0,
				end:     1,
				current: 1,
			},
		},
		{
			in: testin{
				start: 1,
				end:   2,
			},
			out: testout{
				start:   1,
				end:     2,
				current: 2,
			},
		},
		{
			in: testin{
				start: 1,
				end:   5,
			},
			out: testout{
				start:   1,
				end:     5,
				current: 2,
			},
		},
	}

	for _, c := range cases {
		sut := NewRange(c.in.start, c.in.end)

		ok := sut.Next()
		if !ok {
			t.Errorf("want true\tgot %v", ok)
		}

		want := c.out.start
		got := sut.Start()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}

		want = c.out.end
		got = sut.End()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}

		want = c.out.current
		got = sut.Current()
		if want != got {
			t.Errorf("want: %v\tgot: %v", want, got)
		}
	}
}
