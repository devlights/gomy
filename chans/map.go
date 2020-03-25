package chans

type (
	// MapFunc -- chans.Map にて利用されるチャネルの各要素に適用する関数です。
	MapFunc func(interface{}) interface{}

	// MapValue -- chans.Map にて利用されるデータ型です。
	MapValue struct {
		Before interface{} // 元の値
		After  interface{} // 適用後の値
	}
)

func newMapValue(before, after interface{}) *MapValue {
	return &MapValue{
		Before: before,
		After:  after,
	}
}

// Map -- 関数 fn を入力チャネル in の各要素に適用した結果を返すチャネルを生成します。
//
// 戻り値のチャネルから取得できるデータ型は、*chans.MapValue となっています。
//
// 		for m := range chans.Map(done, inCh, fn) {
// 			if v, ok := m.(*chans.MapValue); ok {
// 				// v.Before で元の値、 v.After で適用後の値が取得できる
// 			}
// 		}
//
func Map(done <-chan struct{}, in <-chan interface{}, fn MapFunc) <-chan interface{} {
	out := make(chan interface{})

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
