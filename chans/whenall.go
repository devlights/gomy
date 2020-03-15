package chans

import (
	"sync"
)

// WhenAll -- 指定した１つ以上のチャネルの全てが閉じられたら、閉じるチャネルを返します。
//
// チャネルを一つも渡さずに呼び出すと、既に close 済みのチャネルを返します。
func WhenAll(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		nilCh := make(chan struct{})
		close(nilCh)

		return nilCh
	case 1:
		return channels[0]
	}

	allDone := make(chan struct{})
	go func() {
		defer close(allDone)

		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, v := range channels {
			go func() {
				defer wg.Done()
				<-v
			}()
		}

		wg.Wait()
	}()

	return allDone
}
