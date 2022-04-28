package times_test

import (
	"testing"
	"time"

	"github.com/devlights/gomy/times"
)

func TestParserParse(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		layout string
		out    string
	}{
		{"1", "2022-04-28 16:23:45.876", "yyyy-MM-dd hh:mm:ss.fff", "2022-04-28 16:23:45.876 +0000 UTC"},
		{"2", "2022-04-28 16:23:45.876 +0000", "yyyy-MM-dd hh:mm:ss.fff loc", "2022-04-28 16:23:45.876 +0000 UTC"},
		{"3", "2022-04-28 16:23:45.876 +0900", "yyyy-MM-dd hh:mm:ss.fff loc", "2022-04-28 16:23:45.876 +0900 +0900"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, err := times.Parser(test.in).Parse(test.layout)
			if err != nil {
				t.Error(err)
			}

			if result.String() != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, result)
			}
		})
	}
}

func TestParserYyyyMmDd(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  time.Time
	}{
		{"yyyymmdd", "20220204", time.Date(2022, time.February, 4, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			r, err := times.Parser(test.in).YyyyMmdd()
			if err != nil {
				t.Error(err)
			}

			if r != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, r)
			}
		})
	}
}

func TestParserYyyyMmDdWithHyphen(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  time.Time
	}{
		{"yyyy-mm-dd", "2022-02-04", time.Date(2022, time.February, 4, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			r, err := times.Parser(test.in).YyyyMmDdWithHyphen()
			if err != nil {
				t.Error(err)
			}

			if r != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, r)
			}
		})
	}
}

func TestParserYyyyMmDdWithSlash(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  time.Time
	}{
		{"yyyy-mm-dd", "2022/02/04", time.Date(2022, time.February, 4, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			r, err := times.Parser(test.in).YyyyMmDdWithSlash()
			if err != nil {
				t.Error(err)
			}

			if r != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, r)
			}
		})
	}
}
