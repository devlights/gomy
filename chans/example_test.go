package chans_test

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
)

func ExampleBridge() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	chSeq := make(chan (<-chan int))
	go func() {
		defer close(chSeq)
		chSeq <- chans.GeneratorContext(procCtx, 1, 2, 3)
		chSeq <- chans.GeneratorContext(procCtx, 4, 5, 6)
	}()

	for v := range chans.BridgeContext(procCtx, chSeq) {
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

	numbers := chans.GeneratorContext(procCtx, data...)
	chunks := chans.BufferContext(procCtx, numbers, count)

	for chunk := range chunks {
		fmt.Println(chunk)
	}

	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7]
}

func ExampleChain() {
	// context and channels
	var (
		ctx  = context.Background()
		base = func(pCtx context.Context) context.Context {
			ctx, cxl := context.WithCancel(pCtx)
			go func() {
				defer cxl()
				time.Sleep(100 * time.Millisecond)
				fmt.Println("base")
			}()
			return ctx
		}(ctx)
	)

	chain1 := chans.ChainContext(ctx, base, func(t time.Time) {
		fmt.Println("chain-1")
	})

	chain2 := chans.ChainContext(ctx, chain1, func(t time.Time) {
		fmt.Println("chain-2")
	})

	<-ctxs.WhenAll(ctx, base, chain1, chain2).Done()

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

	numbers := chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5, 6, 7)
	chunks := chans.ChunkContext(procCtx, numbers, 3)

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

	nums1 := chans.GeneratorContext(procCtx, 1, 2, 3)
	nums2 := chans.GeneratorContext(procCtx, 4, 5, 6)

	for v := range chans.ConcatContext(procCtx, nums1, nums2) {
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

func ExampleConvert() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		numbers   = chans.GeneratorContext(procCtx, 1, 2, 3)
		converted = chans.ConvertContext(procCtx, numbers, func(v int) string { return strconv.Itoa(v) })
	)

	for v := range converted {
		fmt.Printf("[%T]%q\n", v, v)
	}

	// Output:
	// [string]"1"
	// [string]"2"
	// [string]"3"
}

func ExampleEnumerate() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	numbers := chans.GeneratorContext(procCtx, 9, 8, 7)
	values := chans.EnumerateContext(procCtx, numbers)

	for v := range values {
		fmt.Printf("%d:%v\n", v.Index, v.Value)
	}

	// Output:
	// 0:9
	// 1:8
	// 2:7
}

func ExampleFanIn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	numStream1 := chans.GeneratorContext(procCtx, 1, 2, 3)
	numStream2 := chans.GeneratorContext(procCtx, 4, 5, 6)

	for v := range chans.FanInContext(procCtx, numStream1, numStream2) {
		fmt.Println(v)
	}

	// Unordered output:
	// 4
	// 1
	// 5
	// 2
	// 3
	// 6
}

func ExampleFanOut() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		nums     = chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5, 6)
		callback = func(v int) { fmt.Println(v) }
	)

	dones := chans.FanOutContext(procCtx, nums, 3, callback)
	<-ctxs.WhenAll(procCtx, dones...).Done()

	// Unordered output:
	// 4
	// 1
	// 2
	// 3
	// 6
	// 5
}

func ExampleFanOutWg() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		nums     = chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5, 6)
		callback = func(v int) { fmt.Println(v) }
	)

	wg := chans.FanOutWgContext(procCtx, nums, 3, callback)
	wg.Wait()

	// Unordered output:
	// 4
	// 1
	// 2
	// 3
	// 6
	// 5
}

func ExampleFilter() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		numbers   = chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5)
		predicate = func(v int) bool {
			return v%2 == 0
		}
	)

	for v := range chans.FilterContext(procCtx, numbers, predicate) {
		fmt.Println(v)
	}

	// Output:
	// 2
	// 4
}

func ExampleForEach() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		in = chans.GeneratorContext(procCtx, 1, 2, 3)
	)

	chans.ForEachContext(procCtx, in, func(v int) {
		fmt.Println(v)
	})

	// Output:
	// 1
	// 2
	// 3
}

func ExampleFromIntCh() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var ints <-chan int = func(pCtx context.Context) <-chan int {
		ctx, cxl := context.WithCancel(pCtx)
		ch := make(chan int)

		go func() {
			defer cxl()
			defer close(ch)
			for i := 0; i < 3; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(procCtx)

	var items <-chan interface{} = chans.FromIntCh(ints)
	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleFromStringCh() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var strs <-chan string = func(pCtx context.Context) <-chan string {
		ctx, cxl := context.WithCancel(pCtx)
		ch := make(chan string)
		go func() {
			defer cxl()
			defer close(ch)
			for _, s := range []string{"h", "e", "l", "l", "o"} {
				select {
				case <-ctx.Done():
				case ch <- s:
				}
			}
		}()
		return ch
	}(procCtx)

	var items <-chan interface{} = chans.FromStringCh(strs)
	for v := range items {
		fmt.Println(v)
	}

	// Output:
	// h
	// e
	// l
	// l
	// o
}

func ExampleGenerator() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	numbers := chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5)
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
		infinite = chans.LoopInfiniteContext(procCtx)
		takes    = chans.TakeContext(procCtx, infinite, 5)
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

func ExampleInterval() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		numbers      = chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5)
		withInterval = chans.IntervalContext(procCtx, numbers, 5*time.Millisecond)
	)

	begin := time.Now()
	for range withInterval {
		// no-op
	}
	elapsed := time.Since(begin)

	fmt.Printf("elapsed <= 35msec: %v\n", elapsed < 50*time.Millisecond)

	// Output:
	// elapsed <= 35msec: true
}

func ExampleLoop() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 10*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	for v := range chans.LoopContext(procCtx, 0, 5) {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleMap() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		numbers = chans.GeneratorContext(procCtx, 1, 2, 3)
		fn      = func(original int) (after int) {
			return original * 2
		}
	)

	for v := range chans.MapContext(procCtx, numbers, fn) {
		fmt.Printf("%v,%v\n", v.Before, v.After)
	}

	// Output:
	// 1,2
	// 2,4
	// 3,6
}

func ExampleMerge() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		inCh1 = chans.GeneratorContext(procCtx, 1, 2, 3)
		inCh2 = chans.GeneratorContext(procCtx, 4, 5, 6)
	)

	for v := range chans.MergeContext(procCtx, inCh1, inCh2) {
		fmt.Println(v)
	}

	// Unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
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

	inCh := chans.GeneratorContext(genCtx, "h", "e", "l", "l", "o")

	var result []interface{}
	for v := range chans.OrDoneContext(procCtx, inCh) {
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

	ch := make(chan int)
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

	repeats := chans.RepeatFnContext(procCtx, func() int { return <-ch })
	takes := chans.TakeContext(procCtx, repeats, 6)

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

	repeats := chans.RepeatContext(procCtx, 1, 2, 3)
	takes := chans.TakeContext(procCtx, repeats, 6)

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

func ExampleSelect() {
	var (
		ch1 = make(chan interface{})
		ch2 = make(chan interface{})
	)
	defer close(ch1)
	defer close(ch2)

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()

	_, v1, _ := chans.Select(ch1, ch2)
	_, v2, _ := chans.Select(ch1, ch2)

	fmt.Println(v1)
	fmt.Println(v2)

	// Unordered output:
	// 1
	// 2
}

func ExampleRecvAny() {
	var (
		ch1 = make(chan interface{})
		ch2 = make(chan interface{})
	)
	defer close(ch1)
	defer close(ch2)

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()

	_, v1, _ := chans.RecvAny(ch1, ch2)
	_, v2, _ := chans.RecvAny(ch1, ch2)

	fmt.Println(v1)
	fmt.Println(v2)

	// Unordered output:
	// 1
	// 2
}

func ExampleRecvAll() {
	var (
		ch1 = make(chan interface{})
		ch2 = make(chan interface{})
	)
	defer close(ch1)
	defer close(ch2)

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()

	for _, v := range chans.RecvAll(ch1, ch2) {
		fmt.Printf("chosen:%d,value:%v\n", v.Chosen, v.Value)
	}

	// Unordered output:
	// chosen:0,value:1
	// chosen:1,value:2
}

func ExampleSkipWhileFn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	numbers := chans.GeneratorContext(procCtx, 1, 1, 1, 4, 5)
	items := chans.SkipWhileFnContext(procCtx, numbers, func() int { return 1 })

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

	numbers := chans.GeneratorContext(procCtx, 1, 1, 1, 4, 5)
	items := chans.SkipWhileContext(procCtx, numbers, 1)

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

	numbers := chans.GeneratorContext(procCtx, 1, 1, 1, 4, 5)
	items := chans.SkipContext(procCtx, numbers, 3)

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

	numbers := chans.GeneratorContext(procCtx, 1, 2, 3, 4, 5)
	takes := chans.TakeContext(procCtx, numbers, 3)

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

	numbers := chans.GeneratorContext(procCtx, 1, 1, 1, 4, 1)
	takes := chans.TakeWhileContext(procCtx, numbers, 1)

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

	numbers := chans.GeneratorContext(procCtx, 1, 1, 1, 4, 1)
	takes := chans.TakeWhileFnContext(procCtx, numbers, func() int { return 1 })

	for v := range takes {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 1
	// 1
}

func ExampleTee() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	numbers := chans.GeneratorContext(procCtx, 1)
	ch1, ch2 := chans.TeeContext(procCtx, numbers)

	var wg sync.WaitGroup
	for _, ch := range []<-chan int{ch1, ch2} {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				fmt.Println(v)
			}
		}(ch)
	}

	wg.Wait()

	// Output:
	// 1
	// 1
}

func ExampleToInt() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		gens <-chan int = chans.GeneratorContext(procCtx, 1, 2)
		ints <-chan int = chans.ToInt(procCtx.Done(), gens, -1)
	)

	for v := range ints {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
}

func ExampleToString() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	var (
		gens <-chan string = chans.GeneratorContext(procCtx, "hello", "world")
		strs <-chan string = chans.ToString(procCtx.Done(), gens, "")
	)

	for v := range strs {
		fmt.Println(v)
	}

	// Output:
	// hello
	// world
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

func ExampleSliceContext() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	var (
		values = []int{1, 2, 3, 4, 5}
		in     = chans.GeneratorContext(ctx, values...)
		out    = chans.SliceContext(ctx, in)
	)

	fmt.Printf("%[1]v (%[1]T)", out)

	// Output:
	// [1 2 3 4 5] ([]int)
}
