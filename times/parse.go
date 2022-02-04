package times

import (
	"time"
)

type (
	Parser string
)

func (me Parser) YyyyMmdd() (time.Time, error) {
	return time.Parse("20060102", string(me))
}

func (me Parser) YyyyMmDdWithHyphen() (time.Time, error) {
	return time.Parse("2006-01-02", string(me))
}

func (me Parser) YyyyMmDdWithSlash() (time.Time, error) {
	return time.Parse("2006/01/02", string(me))
}
