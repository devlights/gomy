package chans_test

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/chans"
)

func ExampleBridge() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	chSeq := make(chan (<-chan interface{}))
	go func() {
		defer close(chSeq)
		chSeq <- chans.Generator(procCtx.Done(), 1, 2, 3)
		chSeq <- chans.Generator(procCtx.Done(), 4, 5, 6)
	}()

	for v := range chans.Bridge(procCtx.Done(), chSeq) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleBuffer() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	var (
		data  = []interface{}{1, 2, 3, 4, 5, 6, 7}
		count = 3
	)

	numbers := chans.Generator(procCtx.Done(), data...)
	chunks := chans.Buffer(procCtx.Done(), numbers, count)

	for chunk := range chunks {
		fmt.Println(chunk)
	}

	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7]
}

func ExampleChain() {
	// functions
	var (
		makeGoroutine = func() <-chan struct{} {
			ch := make(chan struct{})
			go func() {
				defer close(ch)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("base")
			}()
			return ch
		}
	)

	// channels
	var (
		done = make(chan struct{})
		base = makeGoroutine()
	)

	defer close(done)

	chain1 := chans.Chain(done, base, func(t time.Time) {
		fmt.Println("chain-1")
	})

	chain2 := chans.Chain(done, chain1, func(t time.Time) {
		fmt.Println("chain-2")
	})

	<-chans.WhenAll(base, chain1, chain2)

	// Output:
	//
	// base
	// chain-1
	// chain-2
}

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

func ExampleConcat() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	nums1 := chans.Generator(procCtx.Done(), 1, 2, 3)
	nums2 := chans.Generator(procCtx.Done(), 4, 5, 6)

	for v := range chans.Concat(procCtx.Done(), nums1, nums2) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleEnumerate() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 9, 8, 7)
	values := chans.Enumerate(procCtx.Done(), numbers)

	for e := range values {
		if v, ok := e.(*chans.IterValue); ok {
			fmt.Printf("%d:%v\n", v.Index, v.Value)
		}
	}

	// Output:
	// 0:9
	// 1:8
	// 2:7
}

func ExampleForEach() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	for v := range chans.ForEach(procCtx.Done(), 1, 2, 3) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleGenerator() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5)
	for v := range numbers {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleLoopInfinite() {
	// contexts
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 10*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	// channels
	var (
		infinite = chans.LoopInfinite(procCtx.Done())
		takes    = chans.Take(procCtx.Done(), chans.FromIntCh(infinite), 5)
	)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleLoop() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 10*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	for v := range chans.Loop(procCtx.Done(), 0, 5) {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleOrDone() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Minute)
		genCtx, genCxl   = context.WithCancel(mainCtx)
	)

	defer mainCxl()
	defer procCxl()
	defer genCxl()

	inCh := chans.Generator(genCtx.Done(), "h", "e", "l", "l", "o")

	var result []interface{}
	for v := range chans.OrDone(procCtx.Done(), inCh) {
		func() {
			defer procCxl()
			result = append(result, v)
		}()
	}

	fmt.Printf("len(result) <= 2: %v", len(result) <= 2)

	// Output:
	// len(result) <= 2: true
}

func ExampleRepeatFn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			for _, v := range []int{1, 2, 3} {
				select {
				case <-procCtx.Done():
					return
				case ch <- v:
				}
			}
		}
	}()

	repeats := chans.RepeatFn(procCtx.Done(), func() interface{} { return <-ch })
	takes := chans.Take(procCtx.Done(), repeats, 6)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleRepeat() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	repeats := chans.Repeat(procCtx.Done(), 1, 2, 3)
	takes := chans.Take(procCtx.Done(), repeats, 6)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleSkipWhileFn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 1, 1, 4, 5)
	items := chans.SkipWhileFn(procCtx.Done(), numbers, func() interface{} { return 1 })

	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// 4
	// 5
}

func ExampleSkipWhile() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 1, 1, 4, 5)
	items := chans.SkipWhile(procCtx.Done(), numbers, 1)

	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// 4
	// 5
}

func ExampleSkip() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1, 1, 1, 4, 5)
	items := chans.Skip(procCtx.Done(), numbers, 3)

	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// 4
	// 5
}

func ExampleTake() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.ForEach(procCtx.Done(), 1, 2, 3, 4, 5)
	takes := chans.Take(procCtx.Done(), numbers, 3)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleTakeWhile() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.ForEach(procCtx.Done(), 1, 1, 1, 4, 1)
	takes := chans.TakeWhile(procCtx.Done(), numbers, 1)

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 1
	// 1
}

func ExampleTakeWhileFn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.ForEach(procCtx.Done(), 1, 1, 1, 4, 1)
	takes := chans.TakeWhileFn(procCtx.Done(), numbers, func() interface{} { return 1 })

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 1
	// 1
}

func ExampleWhenAll() {
	fn := func(tlimit time.Duration) <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)
			time.Sleep(tlimit)
		}()

		return done
	}

	done1 := fn(100 * time.Millisecond)
	done2 := fn(200 * time.Millisecond)
	done3 := fn(300 * time.Millisecond)

	start := time.Now()
	<-chans.WhenAll(done1, done2, done3)
	elapsed := time.Since(start)

	fmt.Printf("elapsed: about 300msec ==> %v\n", elapsed >= 299*time.Millisecond)

	// Output:
	// elapsed: about 300msec ==> true
}

func ExampleWhenAny() {
	fn := func(tlimit time.Duration) <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)
			time.Sleep(tlimit)
		}()

		return done
	}

	done1 := fn(100 * time.Millisecond)
	done2 := fn(200 * time.Millisecond)
	done3 := fn(300 * time.Millisecond)

	start := time.Now()
	<-chans.WhenAny(done1, done2, done3)
	elapsed := time.Since(start)

	fmt.Printf("elapsed: about 100msec ==> %v\n", elapsed <= 110*time.Millisecond)

	// Output:
	// elapsed: about 100msec ==> true
}
