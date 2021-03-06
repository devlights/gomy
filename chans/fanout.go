package chans

import (
	"sync"
)

// FanOut -- 指定されたチャネルの処理を指定されたワーカーの数でファンアウトします。
//
// チャネルからデータを取得するたびに引数 callback が呼ばれます。
func FanOut(done <-chan struct{}, in <-chan interface{}, workerCount int, callback func(interface{})) []<-chan struct{} {
	outChList := make([]<-chan struct{}, 0)

	for i := 0; i < workerCount; i++ {
		out := make(chan struct{})

		go func() {
			defer close(out)

			for v := range OrDone(done, in) {
				select {
				case <-done:
					return
				default:
				}

				callback(v)
			}
		}()

		outChList = append(outChList, out)
	}

	return outChList
}

// FanOutWg -- FanOut() の sync.WaitGroup を返す版です。
func FanOutWg(done <-chan struct{}, in <-chan interface{}, workerCount int, callback func(interface{})) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range OrDone(done, in) {
				select {
				case <-done:
					return
				default:
				}

				callback(v)
			}
		}()
	}

	return &wg
}
