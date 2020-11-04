package latch

import (
	"sync"
)

type (
	// CountDownLatch は、指定されたカウント数分の非同期処理が完了するまで
	// １つ以上のゴルーチンを待機させる非同期イベントです。
	//
	// Javaの java.util.concurrent.CountDownLatch を参考にしています。
	//     REF: https://docs.oracle.com/javase/jp/8/docs/api/java/util/concurrent/CountDownLatch.html
	//
	// NewCountDownLatch(int) でラッチを生成し、非同期処理側で CountDown() を呼び出します。
	// 生成時に指定したカウント数分の CountDown() 呼び出しが実施されたタイミングでラッチがオープンとなります。
	//
	// 非同期処理の完了を待機する側は Await() を呼び出し、取得したチャネルを監視します。
	// ラッチがオープンとなったタイミングで、このチャネルはクローズされます。
	CountDownLatch struct {
		count       int
		lock        sync.Mutex
		signal      chan struct{}
		latchOpened bool
	}
)

// NewCountDownLatch は、指定したカウント数を使用して *CountDownLatch を生成します。
func NewCountDownLatch(count int) *CountDownLatch {
	return &CountDownLatch{
		count:       count,
		lock:        sync.Mutex{},
		signal:      make(chan struct{}),
		latchOpened: false,
	}
}

// Await は、非同期処理の完了を待機する際に利用できるチャネルを返します。
// このチャネルは、ラッチがオープンした際にクローズされます。
func (c *CountDownLatch) Await() <-chan struct{} {
	defer func() { c.lock.Unlock() }()
	c.lock.Lock()
	c.openLatchIfPossible()
	return c.signal
}

// CountDown は、ラッチをオープンするために必要なカウントを１減らします。
func (c *CountDownLatch) CountDown() {
	defer func() { c.lock.Unlock() }()
	c.lock.Lock()

	select {
	case <-c.signal:
		return
	default:
	}

	c.count--
	c.openLatchIfPossible()
}

func (c *CountDownLatch) openLatchIfPossible() {
	if c.count <= 0 && !c.latchOpened {
		close(c.signal)
		c.latchOpened = true
	}
}
