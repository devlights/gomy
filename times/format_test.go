package times_test

import (
	"testing"
	"time"

	"github.com/devlights/gomy/times"
)

func TestTimeFormat(t *testing.T) {
	type (
		testin struct {
			t  times.Formatter
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

	var (
		f1 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local))
		f2 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local))
		f3 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local))
		f4 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local))
		f5 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, 987654321, time.Local))
	)

	cases := []testcase{
		{
			in: testin{
				t:  f1,
				fn: f1.YyyyMmdd,
			},
			out: testout{
				r: "2999/12/31",
			},
		},
		{
			in: testin{
				t:  f2,
				fn: f2.YyyyMmddHHmmss,
			},
			out: testout{
				r: "2999/12/31 11:12:13",
			},
		},
		{
			in: testin{
				t:  f3,
				fn: f3.YyyyMmddHHmmssWithMilliSec,
			},
			out: testout{
				r: "2999/12/31 11:12:13.987",
			},
		},
		{
			in: testin{
				t:  f4,
				fn: f4.HHmmss,
			},
			out: testout{
				r: "11:12:13",
			},
		},
		{
			in: testin{
				t:  f5,
				fn: f5.HHmmssWithMilliSec,
			},
			out: testout{
				r: "11:12:13.987",
			},
		},
	}

	for _, c := range cases {
		result := c.in.fn(time.Time(c.in.t))
		if c.out.r != result {
			t.Errorf("want: %v\tgot: %v", c.out.r, result)
		}
	}
}
