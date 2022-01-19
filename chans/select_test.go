package chans_test

import (
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestSelect(t *testing.T) {
	tests := []struct {
		name    string
		in, out []interface{}
	}{
		{
			"3-chans",
			[]interface{}{1, 2, 3},
			[]interface{}{3, 2, 1},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			chs := make([]chan interface{}, len(test.in))
			for i, v := range test.in {
				var (
					i, v = i, v
					ch   = make(chan interface{})
				)
				go func() {
					ch <- v
				}()
				chs[i] = ch
			}

			for i := 0; i < len(test.in); i++ {
				chosen, v, ok := chans.Select(chs...)
				if !ok {
					t.Errorf("[recvOk][want] true\t[got] %v", ok)
				}

				if !(chosen >= 0 && chosen < len(test.in)) {
					t.Errorf("[chosen][want] from 0 to %v\t[got] %v", len(test.in)-1, chosen)
				}

				if v == nil {
					t.Errorf("[v][want] not nil\t[got] %v", v)
				}

				found := false
				for _, o := range test.out {
					if v == o {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("[v][want] %v\t[got] %v", test.out, v)
				}
			}
		})
	}
}

func TestRecvAny(t *testing.T) {
	tests := []struct {
		name    string
		in, out []interface{}
	}{
		{
			"3-chans",
			[]interface{}{1, 2, 3},
			[]interface{}{3, 2, 1},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			chs := make([]chan interface{}, len(test.in))
			for i, v := range test.in {
				var (
					i, v = i, v
					ch   = make(chan interface{})
				)
				go func() {
					ch <- v
				}()
				chs[i] = ch
			}

			for i := 0; i < len(test.in); i++ {
				chosen, v, ok := chans.RecvAny(chs...)
				if !ok {
					t.Errorf("[recvOk][want] true\t[got] %v", ok)
				}

				if !(chosen >= 0 && chosen < len(test.in)) {
					t.Errorf("[chosen][want] from 0 to %v\t[got] %v", len(test.in)-1, chosen)
				}

				if v == nil {
					t.Errorf("[v][want] not nil\t[got] %v", v)
				}

				found := false
				for _, o := range test.out {
					if v == o {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("[v][want] %v\t[got] %v", test.out, v)
				}
			}
		})
	}
}
