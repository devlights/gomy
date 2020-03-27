package async

import (
	"sync"

	"github.com/devlights/gomy/chans"
)

type (
	// MergedWaitGroup -- sync.WaitGroup をまとめて管理するための振る舞いを持つインターフェースです。
	MergedWaitGroup interface {
		// Add -- 指定した sync.WaitGroup を管理対象に追加します。
		Add(wg *sync.WaitGroup)
		// Wait -- 内部で管理している sync.WaitGroup が全て完了するまで待機します。
		//
		// 本メソッドは、呼び出すとブロックします。
		Wait()
	}

	mergedWg struct {
		wgs []interface{}
	}
)

// NewMergedWaitGroup -- MergedWaitGroup を生成します。
func NewMergedWaitGroup(wgs ...*sync.WaitGroup) MergedWaitGroup {
	m := new(mergedWg)

	for _, v := range wgs {
		m.wgs = append(m.wgs, v)
	}

	return m
}

// Add -- impl MergedWaitGroup.Add
func (m *mergedWg) Add(wg *sync.WaitGroup) {
	m.wgs = append(m.wgs, wg)
}

// Wait -- impl MergedWaitGroup.Wait
func (m *mergedWg) Wait() {
	done := make(chan struct{})
	defer close(done)

	<-chans.WhenAll(chans.FanOut(done, chans.ForEach(done, m.wgs...), len(m.wgs), func(v interface{}) {
		if wg, ok := v.(*sync.WaitGroup); ok {
			wg.Wait()
		}
	})...)
}
