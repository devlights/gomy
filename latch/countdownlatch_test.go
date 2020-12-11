// Package latch_test は、latch パッケージの外部テストパッケージです。
//
// REF:
//   - https://qiita.com/hogedigo/items/5f491994647aa4a8a905
//   - https://segment.com/blog/5-advanced-testing-techniques-in-go/
package latch_test

import (
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/devlights/gomy/latch"
)

func ExampleCountDownLatch() {
	const (
		latchCount     = 2
		goroutineCount = 4
	)

	var (
		goroutineLog = log.New(os.Stdout, "[goroutine] ", 0)
		waiterLog    = log.New(os.Stdout, "[waiter   ] ", 0)
	)

	var (
		wg sync.WaitGroup
	)

	// make latch
	l := latch.NewCountDownLatch(latchCount)

	// start goroutines
	wg.Add(goroutineCount)
	for i := 0; i < goroutineCount; i++ {
		i := i
		go func() {
			defer func() { wg.Done() }()

			time.Sleep(time.Duration(1+i) * 100 * time.Millisecond)
			goroutineLog.Printf("done [%d]", i)

			l.CountDown()
		}()
	}

	// wait until latch is open
	select {
	case <-l.Await():
		waiterLog.Print("latch opened")
	case <-time.After(goroutineCount * time.Second):
		waiterLog.Print("time over")
	}

	// wait until all goroutine is done
	wg.Wait()

	// Output:
	// [goroutine] done [0]
	// [goroutine] done [1]
	// [waiter   ] latch opened
	// [goroutine] done [2]
	// [goroutine] done [3]
}

func TestCountDownLatch(t *testing.T) {
	cases := []struct {
		name           string
		latchCount     int
		goroutineCount int
		limit          time.Duration
	}{
		{"latch 3, goroutine 3", 3, 3, 2 * time.Second},
		{"latch 2, goroutine 3", 2, 3, 2 * time.Second},
		{"latch 1, goroutine 3", 1, 3, 2 * time.Second},
		{"latch 0, goroutine 3", 0, 3, 2 * time.Second},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var wg sync.WaitGroup
			l := latch.NewCountDownLatch(c.latchCount)

			wg.Add(c.goroutineCount)
			for i := 0; i < c.goroutineCount; i++ {
				i := i
				go func() {
					defer func() { wg.Done() }()
					defer l.CountDown()

					time.Sleep(100 * time.Millisecond)
					t.Logf("[goroutine] countdown [%d]", i)
				}()
			}

			select {
			case <-l.Await():
				t.Log("[await] recv signal")
			case <-time.After(c.limit):
				t.Errorf("[want] no time limit exceeded\t[got] time limited")
			}

			wg.Wait()
		})
	}
}

func TestMultipleWaiters(t *testing.T) {
	cases := []struct {
		name           string
		latchCount     int
		goroutineCount int
		waiterCount    int
		limit          time.Duration
	}{
		{"latch 3, goroutine 3, waiter 3", 3, 3, 3, 2 * time.Second},
		{"latch 2, goroutine 3, waiter 3", 2, 3, 3, 2 * time.Second},
		{"latch 1, goroutine 3, waiter 3", 1, 3, 3, 2 * time.Second},
		{"latch 0, goroutine 3, waiter 3", 0, 3, 3, 2 * time.Second},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var wg sync.WaitGroup
			l := latch.NewCountDownLatch(c.latchCount)

			wg.Add(c.goroutineCount)
			for i := 0; i < c.goroutineCount; i++ {
				i := i
				go func() {
					defer func() { wg.Done() }()
					defer l.CountDown()

					time.Sleep(100 * time.Millisecond)
					t.Logf("[goroutine] countdown [%d]", i)
				}()
			}

			for i := 0; i < c.waiterCount; i++ {
				i := i
				select {
				case <-l.Await():
					t.Logf("[await] recv signal\twaiter:[%d]", i)
				case <-time.After(c.limit):
					t.Errorf("[want] no time limit exceeded\t[got] time limited\twaiter[%d]", i)
				}
			}

			wg.Wait()
		})
	}
}
