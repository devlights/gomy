package times_test

import (
	"testing"
	"time"

	"github.com/devlights/gomy/times"
)

func toNanoSec(milliSec int) int {
	return milliSec * 1000 * 1000
}

func TestFormatterFormat(t *testing.T) {
	tests := []struct {
		name   string
		in     time.Time
		layout string
		out    string
	}{
		{"1", time.Date(2022, time.April, 28, 16, 23, 45, toNanoSec(876), time.UTC), "yyyy-MM-dd", "2022-04-28"},
		{"2", time.Date(2022, time.April, 28, 16, 23, 45, toNanoSec(876), time.UTC), "yyyy-MM-dd hh:mm:ss", "2022-04-28 16:23:45"},
		{"3", time.Date(2022, time.April, 28, 16, 23, 45, toNanoSec(876), time.UTC), "yyyy-MM-dd hh:mm:ss.fff", "2022-04-28 16:23:45.876"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			out := times.Formatter(test.in).Format(test.layout)
			if test.out != out {
				t.Errorf("[want] %v\t[got] %v", test.out, out)
			}
		})
	}
}

func TestTimeFormat(t *testing.T) {
	type (
		testin struct {
			t  times.Formatter
			fn func() string
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
		f1 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, toNanoSec(987), time.Local))
		f2 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, toNanoSec(987), time.Local))
		f3 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, toNanoSec(987), time.Local))
		f4 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, toNanoSec(987), time.Local))
		f5 = times.Formatter(time.Date(2999, 12, 31, 11, 12, 13, toNanoSec(987), time.Local))
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
		c := c

		result := c.in.fn()
		if c.out.r != result {
			t.Errorf("want: %v\tgot: %v", c.out.r, result)
		}
	}
}
