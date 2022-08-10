package times

import (
	"time"
)

type (
	// Formatter -- time.Format() を簡易に実行するための型です.
	Formatter time.Time
)

func convTime(f Formatter) time.Time {
	o := time.Time(f)
	t := time.Date(o.Year(), o.Month(), o.Day(), o.Hour(), o.Minute(), o.Second(), o.Nanosecond(), o.Location())

	return t
}

// Format -- 指定されたレイアウトに従って time.Format() を呼び出します。
//
// 指定できる書式は以下です。
//
//	yyyy: 年 (4桁)
//	MM  : 月 (2桁)
//	dd  : 日 (2桁)
//	hh  : 時 (2桁)(24h)
//	mm  : 分 (2桁)
//	ss  : 秒 (2桁)
//	fff : ミリ秒 (3桁)
//	loc : タイムゾーン (+0900などの部分に対応します)
//
// 例
//
//	2022-04-28 16:23:45.876 +0900
//	yyyy-MM-dd hh:mm:ss.fff loc
func (me Formatter) Format(layout string) string {
	return convTime(me).Format(convLayout(layout))
}

// YyyyMmdd -- 指定された time.Time を yyyy/MM/dd 形式にフォーマットして返します。
func (me Formatter) YyyyMmdd() string {
	return convTime(me).Format("2006/1/2")
}

// YyyyMmddHHmmss -- 指定された time.Time を yyyy/MM/dd HH:mm:ss 形式にフォーマットして返します。
func (me Formatter) YyyyMmddHHmmss() string {
	return convTime(me).Format("2006/1/2 15:04:05")
}

// YyyyMmddHHmmssWithMilliSec -- YyyyMmddHHmmss() の結果にミリ秒を付与したものを返します。
func (me Formatter) YyyyMmddHHmmssWithMilliSec() string {
	return convTime(me).Format("2006/1/2 15:04:05.000")
}

// HHmmss -- 指定された time.Time を HH:mm:ss 形式にフォーマットして返します。
func (me Formatter) HHmmss() string {
	return convTime(me).Format("15:04:05")
}

// HHmmssWithMilliSec -- HHmmss() の結果にミリ秒を付与したものを返します。
func (me Formatter) HHmmssWithMilliSec() string {
	return convTime(me).Format("15:04:05.000")
}
