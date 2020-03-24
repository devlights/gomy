package times

import (
	"time"
)

// YyyyMmdd -- 指定された time.Time を yyyy/MM/dd 形式にフォーマットして返します。
func YyyyMmdd(t time.Time) string {
	return t.Format("2006/1/2")
}

// YyyyMmddHHmmss -- 指定された time.Time を yyyy/MM/dd HH:mm:ss 形式にフォーマットして返します。
func YyyyMmddHHmmss(t time.Time) string {
	return t.Format("2006/1/2 15:04:05")
}

// YyyyMmddHHmmssWithMilliSec -- YyyyMmddHHmmss() の結果にミリ秒を付与したものを返します。
func YyyyMmddHHmmssWithMilliSec(t time.Time) string {
	return t.Format("2006/1/2 15:04:05.000")
}

// HHmmss -- 指定された time.Time を HH:mm:ss 形式にフォーマットして返します。
func HHmmss(t time.Time) string {
	return t.Format("15:04:05")
}

// HHmmssWithMilliSec -- HHmmss() の結果にミリ秒を付与したものを返します。
func HHmmssWithMilliSec(t time.Time) string {
	return t.Format("15:04:05.000")
}
