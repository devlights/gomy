package times

import (
	"strings"
	"time"
)

type (
	// Parse -- time.Parse() を簡易に実行するための型です.
	Parser string
)

func convLayout(layout string) string {
	var (
		l = layout
	)

	l = strings.Replace(l, "yyyy", "2006", 1)
	l = strings.Replace(l, "MM", "01", 1)
	l = strings.Replace(l, "dd", "02", 1)
	l = strings.Replace(l, "hh", "15", 1)
	l = strings.Replace(l, "mm", "04", 1)
	l = strings.Replace(l, "ss", "05", 1)
	l = strings.Replace(l, "fff", "000", 1)
	l = strings.Replace(l, "loc", "-0700", 1)

	return l
}

// Parse -- 指定されたレイアウトに従って time.Parse() を呼び出します。
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
func (me Parser) Parse(layout string) (time.Time, error) {
	t, err := time.Parse(convLayout(layout), string(me))
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (me Parser) YyyyMmdd() (time.Time, error) {
	return time.Parse("20060102", string(me))
}

func (me Parser) YyyyMmDdWithHyphen() (time.Time, error) {
	return time.Parse("2006-01-02", string(me))
}

func (me Parser) YyyyMmDdWithSlash() (time.Time, error) {
	return time.Parse("2006/01/02", string(me))
}
