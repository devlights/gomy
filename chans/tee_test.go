package chans_test

import (
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestTee(t *testing.T) {
	type (
		testin struct {
			data []interface{}
		}
		testout struct {
			result1 []interface{}
			result2 []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				data: []interface{}{
					"hello",
					"world",
				},
			},
			out: testout{
				result1: []interface{}{
					"hello",
					"world",
				},
				result2: []interface{}{
					"hello",
					"world",
				},
			},
		},
	}

	for _, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			inCh := make(chan interface{}, len(c.in.data))
			func() {
				defer close(inCh)
				for _, v := range c.in.data {
					inCh <- v
				}
			}()

			out1, out2 := chans.Tee(done, inCh)

			result1 := make([]interface{}, 0)
			result2 := make([]interface{}, 0)
			for v := range out1 {
				v2 := <-out2

				t.Logf("out1: %v\tout2: %v", v, v2)

				result1 = append(result1, v)
				result2 = append(result2, v2)
			}

			if len(c.in.data) != len(result1) {
				t.Errorf("[len(c.in.data) != len(result1)] want: %d\tgot: %d", len(c.in.data), len(result1))
			}

			if len(c.in.data) != len(result2) {
				t.Errorf("[len(c.in.data) != len(result2)] want: %d\tgot: %d", len(c.in.data), len(result2))
			}

			if len(result1) != len(result2) {
				t.Errorf("[len(result1) != len(result2)] want: %d\tgot: (%d, %d)", len(c.in.data), len(result1), len(result2))
			}
		}()
	}
}
