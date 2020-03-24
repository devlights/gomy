package times

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	type (
		testin struct {
			t  time.Time
			fn func(t time.Time) string
		}
		testout struct {
			r string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				t:  time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local),
				fn: YyyyMmdd,
			},
			out: testout{
				r: "2999/12/31",
			},
		},
		{
			in: testin{
				t:  time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local),
				fn: YyyyMmddHHmmss,
			},
			out: testout{
				r: "2999/12/31 11:12:13",
			},
		},
		{
			in: testin{
				t:  time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local),
				fn: YyyyMmddHHmmssWithMilliSec,
			},
			out: testout{
				r: "2999/12/31 11:12:13.987",
			},
		},
		{
			in: testin{
				t:  time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local),
				fn: HHmmss,
			},
			out: testout{
				r: "11:12:13",
			},
		},
		{
			in: testin{
				t:  time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local),
				fn: HHmmssWithMilliSec,
			},
			out: testout{
				r: "11:12:13.987",
			},
		},
	}

	for _, c := range cases {
		result := c.in.fn(c.in.t)
		if c.out.r != result {
			t.Errorf("want: %v\tgot: %v", c.out.r, result)
		}
	}
}
