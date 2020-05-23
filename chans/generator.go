package chans

// Generator -- 指定されたデータを出力するチャネルを生成します。
//
// ForEach関数のエイリアスです。
func Generator(done <-chan struct{}, in ...interface{}) <-chan interface{} {
	return ForEach(done, in...)
}
