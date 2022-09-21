package chans_test

import (
	"context"
	"testing"

	"github.com/devlights/gomy/chans"
	"golang.org/x/exp/slices"
)

func TestMerge(t *testing.T) {
	// Arrange
	var (
		ctx     = context.Background()
		values1 = []int{1, 2, 3}
		values2 = []int{4, 5, 6}
		in1     = chans.GeneratorContext(ctx, values1...)
		in2     = chans.GeneratorContext(ctx, values2...)
		out     = []int{1, 2, 3, 4, 5, 6}
	)

	// Act
	var ch <-chan int = chans.MergeContext(ctx, in1, in2)
	chans.ForEachContext(ctx, ch, func(v int) {
		// Assert
		if !slices.Contains(out, v) {
			t.Errorf("%v does not contains in %v", v, out)
		}
	})
}
