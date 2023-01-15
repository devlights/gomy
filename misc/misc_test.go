package misc

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	result := make([]int, 0, 4)
	for v := range PrimeNumber(context.Background(), 10) {
		result = append(result, v)
	}

	if len(result) != 4 {
		t.Errorf("[want] 4\t[got] %v", len(result))
	}

	if !reflect.DeepEqual([]int{2, 3, 5, 7}, result) {
		t.Errorf("[want] 2,3,5,7\t[got] %v", result)
	}
}

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
