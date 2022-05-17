package errs

import (
	"errors"
	"testing"
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

	v := Forgot(fn1())
	if v != "hello" {
		t.Errorf("[want] %s\t[got] %s", "hello", v)
	}

	v2 := Forgot(fn2())
	if v2 != 100 {
		t.Errorf("[want] %d\t[got] %d", 100, v2)
	}
}
