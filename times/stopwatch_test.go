package times_test

import (
	"testing"
	"time"

	"github.com/devlights/gomy/times"
)

func TestStopwatch(t *testing.T) {
	// Arrange
	fn := func(start time.Time) {
		time.Sleep(100 * time.Millisecond)
	}

	// Act
	elapsed := times.Stopwatch(fn)

	// Assert
	if elapsed < 1*time.Millisecond {
		t.Errorf("[want] greater than 1sec\t[got] %v", elapsed)
	}
}
