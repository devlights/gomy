package chans

import "reflect"

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
