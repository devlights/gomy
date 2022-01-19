package chans_test

import (
	"context"
	"fmt"
	"sync"
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

func ExampleFanIn() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numStream1 := chans.Generator(procCtx.Done(), 1, 2, 3)
	numStream2 := chans.Generator(procCtx.Done(), 4, 5, 6)

	for v := range chans.FanIn(procCtx.Done(), numStream1, numStream2) {
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
		nums     = chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5, 6)
		callback = func(v interface{}) { fmt.Println(v) }
	)

	dones := chans.FanOut(procCtx.Done(), nums, 3, callback)
	<-chans.WhenAll(dones...)

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
		nums     = chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5, 6)
		callback = func(v interface{}) { fmt.Println(v) }
	)

	wg := chans.FanOutWg(procCtx.Done(), nums, 3, callback)
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
		numbers   = chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5)
		predicate = func(v interface{}) bool {
			if i, ok := v.(int); ok {
				return i%2 == 0
			}
			return false
		}
	)

	for v := range chans.Filter(procCtx.Done(), numbers, predicate) {
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

	for v := range chans.ForEach(procCtx.Done(), 1, 2, 3) {
		fmt.Println(v)
	}

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

func ExampleInterval() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	var (
		numbers      = chans.Generator(procCtx.Done(), 1, 2, 3, 4, 5)
		withInterval = chans.Interval(procCtx.Done(), numbers, 5*time.Millisecond)
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

func ExampleMap() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	var (
		numbers = chans.Generator(procCtx.Done(), 1, 2, 3)
		fn      = func(original interface{}) (after interface{}) {
			if i, ok := original.(int); ok {
				return i * 2
			}
			return nil
		}
	)

	for v := range chans.Map(procCtx.Done(), numbers, fn) {
		if m, ok := v.(*chans.MapValue); ok {
			fmt.Printf("%v,%v\n", m.Before, m.After)
		}
	}

	// Output:
	// 1,2
	// 2,4
	// 3,6
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

func ExampleTee() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 50*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	numbers := chans.Generator(procCtx.Done(), 1)
	ch1, ch2 := chans.Tee(procCtx.Done(), numbers)

	var wg sync.WaitGroup
	for _, ch := range []<-chan interface{}{ch1, ch2} {
		wg.Add(1)
		go func(ch <-chan interface{}) {
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
		gens <-chan interface{} = chans.Generator(procCtx.Done(), 1, 2)
		ints <-chan int         = chans.ToInt(procCtx.Done(), gens, -1)
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
		gens <-chan interface{} = chans.Generator(procCtx.Done(), "hello", "world")
		strs <-chan string      = chans.ToString(procCtx.Done(), gens, "")
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
