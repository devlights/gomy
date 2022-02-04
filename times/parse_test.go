package times_test

import (
	"testing"
	"time"

	"github.com/devlights/gomy/times"
)

func TestParserYyyyMmDd(t *testing.T) {
	tests := []struct{
		name string
		in string
		out time.Time
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
	tests := []struct{
		name string
		in string
		out time.Time
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
	tests := []struct{
		name string
		in string
		out time.Time
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