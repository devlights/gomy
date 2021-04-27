package signals

type (
	// Ready は、処理が始まったことを他のゴルーチンに知らせたい場合に利用できる簡易イベントです.
	//
	// Ready は、一度 Signal すると完了状態から戻ることはありません。
	// 何度も、待機と通知を繰り返したい場合は *sync.Cond の利用を検討してください。
	Ready struct {
		*ready
	}

	ready struct {
		ch chan struct{}
	}
)

// NewReady は、新しい *signals.Ready を生成して返します.
func NewReady() *Ready {
	r := new(ready)
	r.ch = make(chan struct{})

	return &Ready{r}
}

// Signal は、処理を開始するための準備が整ったことを通知します.
func (me *Ready) Signal() {
	close(me.ch)
}

// Wait は、他のゴルーチンが Signal するまで待ちます.
//
// 本メソッドは、Signal メソッドがコールされるまで呼び出し元をブロックします。
//
// 一度 Signal された以降に本メソッドを呼び出すと即座に処理を返します。
func (me *Ready) Wait() {
	<-me.ch
}
