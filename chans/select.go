package chans

import (
	"reflect"
	"sync"
)

type (
	// SelectValue -- chans.RecvAll() で利用されるデータ型です.
	SelectValue struct {
		Chosen int         // 選択されたチャネルのインデックス
		Value  interface{} // 受信した値
	}
)

// Eq -- 同じデータかどうかを判定します.
func (me SelectValue) Eq(other SelectValue) bool {
	if me == other {
		return true
	}

	if me.Chosen == other.Chosen && me.Value == other.Value {
		return true
	}

	return false
}

// Select -- 指定されたチャネルリストから一つ値を取得します。どのチャネルが選択されるかは非決定的です。
//
// 内部で reflect.Select() を呼び出しており、戻り値はそれに準じています。
//
// REFERENCES:
//   - https://pkg.go.dev/reflect#Select
//   - https://pkg.go.dev/reflect#SelectCase
//   - https://dev.to/hgsgtk/handling-with-arbitrary-channels-by-reflectselect-4d5g
func Select(chs ...chan interface{}) (chosen int, v interface{}, ok bool) {
	if len(chs) == 0 {
		return -1, nil, false
	}

	scs := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		sc := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
		scs[i] = sc
	}

	chosen, recv, ok := reflect.Select(scs)
	if ok {
		v = recv.Interface()
	}

	return chosen, v, ok
}

// RecvAny -- 指定されたチャネルリストから一つ値を取得します。どのチャネルが選択されるかは非決定的です。
//
// See: chans.Select
func RecvAny(chs ...chan interface{}) (chosen int, v interface{}, ok bool) {
	return Select(chs...)
}

// RecvAll -- 指定されたチャネルリストの全てから１つ値を取得して返却します。
func RecvAll(chs ...chan interface{}) []SelectValue {
	var (
		wg  sync.WaitGroup
		ret = make([]SelectValue, len(chs))
	)

	wg.Add(len(chs))
	for i := 0; i < len(chs); i++ {
		go func(i int) {
			defer wg.Done()

			if chosen, v, ok := Select(chs...); ok {
				ret[i] = SelectValue{chosen, v}
			}
		}(i)
	}

	wg.Wait()
	return ret
}
