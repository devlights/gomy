package signals_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/devlights/gomy/signals"
)

func ExampleReady() {
	// 以下は処理の開始を待つゴルーチンが複数存在する状態で
	// 値を提供するゴルーチンが自身の準備が整った後 *signals.Ready を用いて
	// 待機しているゴルーチンに開始許可を通知する例です.

	var (
		wg sync.WaitGroup // 消費者側ゴルーチンの完了待機用
	)

	// Producer
	ready, values := func() (*signals.Ready, <-chan int) {
		ready := signals.NewReady()
		values := make(chan int)

		go func() {
			defer close(values)

			// 開始まで時間が掛かる処理をシミュレート
			time.Sleep(2 * time.Second)
			// 待機している他のゴルーチンに開始許可を通知
			ready.Signal()

			for _, v := range []int{1, 2, 3, 4, 5} {
				values <- v

				// 各ゴルーチンが値を受け取れるように意図的に少し間を空ける
				time.Sleep(20 * time.Millisecond)
			}
		}()

		return ready, values
	}()

	// Consumer
	for _, i := range []int{1, 2, 3} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 開始許可が出るまで待機
			fmt.Printf("[%d]:待機開始\n", i)
			ready.Wait()
			fmt.Printf("[%d]:待機解除\n", i)

			for v := range values {
				fmt.Println(v)
			}
		}(i)
	}

	wg.Wait()

	// Unordered output:
	// [1]:待機開始
	// [2]:待機開始
	// [3]:待機開始
	// [1]:待機解除
	// [2]:待機解除
	// [3]:待機解除
	// 1
	// 2
	// 3
	// 4
	// 5
}
