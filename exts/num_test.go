package exts_test

import (
	"strings"
	"testing"

	"github.com/devlights/gomy/exts"
)

func TestNumToStr(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  string
	}{
		{"1", 1, "1"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			n := exts.Num(test.in)
			v := n.ToStr()

			if v != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, v)
			}
		})
	}
}

func TestNumTimes(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  string
	}{
		{"3", 3, "0,1,2"},
		{"0", 0, ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			n := exts.Num(test.in)
			l := make([]string, 0)

			n.Times(func(i int) {
				l = append(l, exts.Num(i).ToStr())
			})
			s := strings.Join(l, ",")

			if s != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, s)
			}
		})
	}
}

func TestNumUpto(t *testing.T) {
	tests := []struct {
		name   string
		in, to int
		out    string
	}{
		{"3-5", 3, 5, "3,4,5"},
		{"5-3", 5, 3, ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			n := exts.Num(test.in)
			l := make([]string, 0)

			n.Upto(test.to, func(i int) {
				l = append(l, exts.Num(i).ToStr())
			})
			s := strings.Join(l, ",")

			if s != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, s)
			}
		})
	}
}

func TestNumDownto(t *testing.T) {
	tests := []struct {
		name   string
		in, to int
		out    string
	}{
		{"5-3", 5, 3, "5,4,3"},
		{"3-5", 3, 5, ""},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			n := exts.Num(test.in)
			l := make([]string, 0)

			n.Downto(test.to, func(i int) {
				l = append(l, exts.Num(i).ToStr())
			})
			s := strings.Join(l, ",")

			if s != test.out {
				t.Errorf("[want] %v\t[got] %v", test.out, s)
			}
		})
	}
}
