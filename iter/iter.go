package iter

// Range は、指定された回数分ループ可能な空スライスを生成して返します。
//
// 元ネタは https://github.com/bradfitz/iter/blob/master/iter.go です。
func Range(n int) []struct{} {
	return make([]struct{}, n)
}

// RangeFn は、Range に処理関数を指定できるバージョンです。挙動は同じです。
//
// 処理中にエラーが発生した場合、内部ループはそこで停止し
// (エラーが発生したインデックス, エラー) を返します。
func RangeFn(n int, fn func(i int) error) (int, error) {
	for i := range Range(n) {
		err := fn(i)
		if err != nil {
			return i, err
		}
	}

	return -1, nil
}