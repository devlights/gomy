package chans_test

import (
	"context"
	"testing"

	"slices"

	"github.com/devlights/gomy/chans"
)

func TestForEachContext(t *testing.T) {
	// Arrange
	var (
		ctx    = context.Background()
		values = []int{1, 2, 3, 4, 5}
		in     = chans.GeneratorContext(ctx, values...)
		out    = make([]int, 0)
	)

	// Act
	chans.ForEachContext(ctx, in, func(v int) {
		out = append(out, v)
	})

	// Assert
	if !slices.Equal(values, out) {
		t.Errorf("[want] %v\t[got] %v", values, out)
	}
}

func TestForEach(t *testing.T) {
	// Arrange
	var (
		done   = make(chan struct{})
		values = []int{1, 2, 3, 4, 5}
		in     = chans.Generator(done, values...)
		out    = make([]int, 0)
	)

	// Act
	chans.ForEach(done, in, func(v int) {
		out = append(out, v)
	})

	// Assert
	if !slices.Equal(values, out) {
		t.Errorf("[want] %v\t[got] %v", values, out)
	}
}
