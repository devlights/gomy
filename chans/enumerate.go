package chans

type (
	// IterValue -- chans.Enumerate() にて利用されるデータ型です。
	IterValue[T any] struct {
		Index int // インデックス
		Value T   // 値
	}
)

func newIterValue[T any](i int, v T) *IterValue[T] {
	return &IterValue[T]{
		Index: i,
		Value: v,
	}
}

// Enumerate -- 指定された入力チャネルの要素に対してインデックスを付与したデータを返すチャネルを生成します。
//
// 戻り値のチャネルから取得できるデータ型は、*chans.IterValue となっています。
//
// 		for v := range chans.Enumerate(done, inCh) {
// 			// v.Index でインデックス、 v.Value で値が取得できる
// 		}
//
func Enumerate[T any](done <-chan struct{}, in <-chan T) <-chan *IterValue[T] {
	out := make(chan *IterValue[T])

	go func() {
		defer close(out)

		index := 0

	ChLoop:
		for {
			select {
			case <-done:
				break ChLoop
			case v, ok := <-in:
				if !ok {
					break ChLoop
				}

				select {
				case out <- newIterValue(index, v):
					index++
				case <-done:
				}
			}
		}
	}()

	return out
}
