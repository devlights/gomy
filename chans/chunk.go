package chans

// Chunkは、入力を指定した件数分に束ねてデータを返すチャネルを生成します.
//
// Buffer関数のエイリアスです。
func Chunk(done <-chan struct{}, in <-chan interface{}, count int) <-chan []interface{} {
	return Buffer(done, in, count)
}
