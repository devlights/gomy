package chans_test

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleChunk() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5, 6, 7)
	chunks := chans.Chunk(procCtx.Done(), numbers, 3)

	for chunk := range chunks {
		fmt.Println(chunk)
	}

	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7]
}
