package promise

import (
	"fmt"
	"sync"
	"time"
)

func ExamplePromise() {
	var (
		f1, p1 = NewPromise[int]()
		f2, p2 = NewPromise[string]()
		wg     sync.WaitGroup
	)

	wg.Add(3)

	// f1に依存する処理
	//
	// f2の状態に関係なく、f1が完了したら完了する
	go func(f *Future[int]) {
		defer wg.Done()
		fmt.Printf("f1=%v\n", f1.Get())
	}(f1)

	// f2に依存する処理
	//
	// f1の状態に関係なく、f2が完了したら完了する
	go func(f *Future[string]) {
		defer wg.Done()
		fmt.Printf("f2=%v\n", f2.Get())
	}(f2)

	// f1とf2に依存する処理
	//
	// f1, f2の２つのFutureが完了しないと処理が完了しない
	go func(f1 *Future[int], f2 *Future[string]) {
		defer wg.Done()
		fmt.Printf("f1=%v\tf2=%v\n", f1.Get(), f2.Get())
	}(f1, f2)

	// 100ms後にf1に値を提供する
	//
	// このPromiseが値を提供するまで対応するFuture(f1)は結果を返さない
	go func(p *Promise[int]) {
		time.Sleep(100 * time.Millisecond)
		p.Submit(999)
	}(p1)

	// 500ms後にf2に値を提供する
	//
	// このPromiseが値を提供するまで対応するFuture(f2)は結果を返さない
	go func(p *Promise[string]) {
		time.Sleep(500 * time.Millisecond)
		p.Submit("hello world")
	}(p2)

	wg.Wait()

	// Unordered output:
	// f1=999
	// f2=hello world
	// f1=999	f2=hello world
}
