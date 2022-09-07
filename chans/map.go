package chans

type (
	// MapFunc -- chans.Map にて利用されるチャネルの各要素に適用する関数です。
	MapFunc[T any, R any] func(T) R

	// MapValue -- chans.Map にて利用されるデータ型です。
	MapValue[T any, R any] struct {
		Before T // 元の値
		After  R // 適用後の値
	}
)

func newMapValue[T any, R any](before T, after R) *MapValue[T, R] {
	return &MapValue[T, R]{
		Before: before,
		After:  after,
	}
}

// Map -- 関数 fn を入力チャネル in の各要素に適用した結果を返すチャネルを生成します。
//
// 戻り値のチャネルから取得できるデータ型は、*chans.MapValue となっています。
//
//	for v := range chans.Map(done, inCh, fn) {
//		// v.Before で元の値、 v.After で適用後の値が取得できる
//	}
func Map[T any, R any](done <-chan struct{}, in <-chan T, fn MapFunc[T, R]) <-chan *MapValue[T, R] {
	out := make(chan *MapValue[T, R])

	go func() {
		defer close(out)

	ChLoop:
		for {
			select {
			case <-done:
				break ChLoop
			case v, ok := <-in:
				if !ok {
					break ChLoop
				}

				before := v
				after := fn(v)

				select {
				case out <- newMapValue(before, after):
				case <-done:
				}
			}
		}
	}()

	return out
}
