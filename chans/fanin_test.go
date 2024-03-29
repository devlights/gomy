package chans_test

import (
	"context"
	"testing"

	"github.com/devlights/gomy/chans"
	"golang.org/x/exp/slices"
)

func TestFanInContext(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3, 4}
		ch1    = chans.Generator(ctx.Done(), values[0], values[1])
		ch2    = chans.Generator(ctx.Done(), values[2], values[3])
	)

	// Act
	var ret <-chan int = chans.FanInContext(ctx, ch1, ch2)

	// Assert
	for v := range ret {
		if !slices.Contains(values, v) {
			t.Errorf("value %v is not included in %v", v, values)
		}
	}
}

func TestFanIn(t *testing.T) {
	type (
		testin struct {
			data [][]interface{}
		}
		testout struct {
			result []interface{}
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{data: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
			}},
			out: testout{result: []interface{}{
				1, 2, 3, 4, 5, 6,
			}},
		},
	}

	for _, c := range cases {
		func() {
			done := make(chan struct{})
			defer close(done)

			chList := make([]<-chan interface{}, 0, len(c.in.data))
			for _, list := range c.in.data {
				func() {
					ch := make(chan interface{}, len(list))
					defer close(ch)

					for _, v := range list {
						ch <- v
					}

					chList = append(chList, ch)
				}()
			}

			fanInCh := chans.FanIn(done, chList...)

			results := make([]interface{}, 0)
			for v := range fanInCh {
				t.Log(v)
				results = append(results, v)
			}

			// fan-in の場合は、flatten と異なり取得順序は不定となるので個数でテストする
			t.Logf("[c.out.result] %v", c.out.result)
			t.Logf("[results     ] %v", results)

			if len(c.out.result) != len(results) {
				t.Errorf("want: %v\tgot: %v", c.out.result, results)
			}
		}()
	}
}
