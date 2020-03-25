package chans

type (
	// IterValue -- chans.Enumerate() にて利用されるデータ型です。
	IterValue struct {
		Index int         // インデックス
		Value interface{} // 値
	}
)

func newIterValue(i int, v interface{}) *IterValue {
	return &IterValue{
		Index: i,
		Value: v,
	}
}

// Enumerate -- 指定された入力チャネルの要素に対してインデックスを付与したデータを返すチャネルを生成します。
//
// 戻り値のチャネルから取得できるデータ型は、*chans.IterValue となっています。
//
// 		for e := range chans.Enumerate(done, inCh) {
// 			if v, ok := e.(*IterValue); ok {
// 				// v.Index でインデックス、 v.Value で値が取得できる
// 			}
// 		}
//
func Enumerate(done <-chan struct{}, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

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
