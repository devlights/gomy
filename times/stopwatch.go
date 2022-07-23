package times

import "time"

// Stopwatch -- fnの実行に掛かった所要時間を返します。
func Stopwatch(fn func(start time.Time)) time.Duration {
	s := time.Now()
	fn(s)
	return time.Since(s)
}
