package chans

import (
	"testing"
	"time"
)

func TestOrDone(t *testing.T) {
	var (
		done    = make(chan struct{})
		inCh    = make(chan interface{})
		chList  = make([]<-chan struct{}, 0, 0)
		results = make([]interface{}, 0, 0)
	)

	chList = append(chList, done)

	// 3 秒後に done チャネルを閉じる
	chList = append(chList, func() <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer close(terminated)
			<-time.After(3 * time.Second)
			close(done)
		}()
		return terminated
	}())

	// 1 秒毎に inCh にデータを送る
	chList = append(chList, func() <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer close(terminated)
			defer close(inCh)

			for v := range Repeat(done, 1) {
				inCh <- v
				<-time.After(1 * time.Second)
			}
		}()
		return terminated
	}())

	// inCh からのデータを出力
	for v := range OrDone(done, inCh) {
		t.Logf("[result] %v", v)
		results = append(results, v)
	}

	<-WhenAll(chList...)

	if len(results) != 3 {
		t.Errorf("want: 3\tgot: %v", len(results))
	}
}
