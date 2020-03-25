package chans

import (
	"math"

	"github.com/devlights/gomy/enumerable"
)

// Loop -- 指定された開始と終了の間、データを返し続けるチャネルを生成します。
func Loop(done <-chan struct{}, start, end int) <-chan int {
	out := make(chan int)

	go func(done <-chan struct{}, start, end int) {
		defer close(out)

		fn := func() func() interface{} {
			r := enumerable.NewRange(start, end)
			return func() interface{} {
				defer r.Next()
				return r.Current()
			}
		}()

		repeatCh := RepeatFn(done, fn)
		takeCh := Take(done, repeatCh, end-start)

		for v := range takeCh {
			if i, ok := v.(int); ok {
				out <- i
			}
		}
	}(done, start, end)

	return out
}

// LoopInfinite -- 無限にループして、データを返し続けるチャネルを生成します。
func LoopInfinite(done <-chan struct{}) <-chan int {
	return Loop(done, 0, math.MaxInt64)
}
