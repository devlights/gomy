package errs_test

import (
	"errors"
	"testing"

	"github.com/devlights/gomy/errs"
)

func TestForgot(t *testing.T) {
	var (
		fn1 = func() (string, error) {
			return "hello", errors.New("this is test")
		}
		fn2 = func() (int, error) {
			return 100, errors.New("this is test 2")
		}
	)

	v := errs.Forgot(fn1())
	if v != "hello" {
		t.Errorf("[want] %s\t[got] %s", "hello", v)
	}

	v2 := errs.Forgot(fn2())
	if v2 != 100 {
		t.Errorf("[want] %d\t[got] %d", 100, v2)
	}
}
