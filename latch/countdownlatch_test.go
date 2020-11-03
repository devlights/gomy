package latch

import (
	"log"
	"os"
	"testing"
	"time"
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

	// make latch
	l := NewCountDownLatch(latchCount)

	// start goroutines
	done := make(chan struct{}, goroutineCount)
	for i := 0; i < goroutineCount; i++ {
		i := i
		go func() {
			defer func() { done <- struct{}{} }()

			time.Sleep(time.Duration(1+i) * time.Second)
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
	for i := 0; i < goroutineCount; i++ {
		<-done
	}

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
			l := NewCountDownLatch(c.latchCount)

			done := make(chan struct{}, c.goroutineCount)
			for i := 0; i < c.goroutineCount; i++ {
				i := i
				go func() {
					defer func() { done <- struct{}{} }()
					defer l.CountDown()

					time.Sleep(1 * time.Second)
					t.Logf("[goroutine] countdown [%d]", i)
				}()
			}

			select {
			case <-l.Await():
				t.Log("[await] recv signal")
			case <-time.After(c.limit):
				t.Errorf("[want] no time limit exceeded\t[got] time limited")
			}

			for i := 0; i < c.goroutineCount; i++ {
				<-done
			}
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
			l := NewCountDownLatch(c.latchCount)

			done := make(chan struct{}, c.goroutineCount)
			for i := 0; i < c.goroutineCount; i++ {
				i := i
				go func() {
					defer func() { done <- struct{}{} }()
					defer l.CountDown()

					time.Sleep(1 * time.Second)
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

			for i := 0; i < c.goroutineCount; i++ {
				<-done
			}
		})
	}
}
