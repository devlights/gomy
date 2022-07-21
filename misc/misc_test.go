package misc

import (
	"errors"
	"testing"
)

func TestOpenBrowserUrlEmpty(t *testing.T) {
	var (
		url = ""
	)

	err := OpenBrowser(url)
	if err == nil {
		t.Error("[want] error\t[got] nil")
	}

	if !errors.Is(err, ErrURLIsEmpty) {
		t.Errorf("[want] UrlIsEmpty\t[got] %v", err)
	}

	t.Log(err)
}
