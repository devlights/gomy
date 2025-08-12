package chans_test

import (
	"context"
	"testing"

	"slices"

	"github.com/devlights/gomy/chans"
)

func TestSliceContext(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3, 4, 5}
		in     = chans.GeneratorContext(ctx, values...)
	)

	// Act
	var ret []int = chans.SliceContext(ctx, in)

	// Assert
	if !slices.Equal(values, ret) {
		t.Errorf("[want] %v\t[got] %v", values, ret)
	}
}

func TestSlice(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3, 4, 5}
		in     = chans.GeneratorContext(ctx, values...)
	)

	// Act
	var ret []int = chans.Slice(ctx.Done(), in)

	// Assert
	if !slices.Equal(values, ret) {
		t.Errorf("[want] %v\t[got] %v", values, ret)
	}
}
