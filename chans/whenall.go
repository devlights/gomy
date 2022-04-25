package chans

import (
	"sync"
)

// WhenAll -- 指定した１つ以上のチャネルの全てが閉じられたら、閉じるチャネルを返します。
//
// チャネルを一つも渡さずに呼び出すと、既に close 済みのチャネルを返します。
func WhenAll[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		nilCh := make(chan T)
		close(nilCh)

		return nilCh
	case 1:
		return channels[0]
	}

	allDone := make(chan T)
	go func() {
		defer close(allDone)

		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, v := range channels {
			go func(ch <-chan T) {
				defer wg.Done()
				<-ch
			}(v)
		}

		wg.Wait()
	}()

	return allDone
}
