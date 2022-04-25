package chans

// Chunkは、入力を指定した件数分に束ねてデータを返すチャネルを生成します.
//
// Buffer関数のエイリアスです。
func Chunk[T any](done <-chan struct{}, in <-chan T, count int) <-chan []T {
	return Buffer(done, in, count)
}
