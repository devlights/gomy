package chans

// Generator -- 指定されたデータを出力するチャネルを生成します。
//
// ForEach関数のエイリアスです。
func Generator[T any](done <-chan struct{}, in ...T) <-chan T {
	return ForEach(done, in...)
}
