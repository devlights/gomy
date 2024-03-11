package promise

type (
	// Promise[T] は、「将来値を提供するという約束」を表します。
	Promise[T any] struct {
		f *Future[T]
	}
)

func (me *Promise[T]) Submit(v T) {
	me.f.set(v)
}

// NewPromise は、FutureとPromiseを生成して返します。
func NewPromise[T any]() (*Future[T], *Promise[T]) {
	f := &Future[T]{wait: make(chan struct{})}
	p := &Promise[T]{f: f}

	return f, p
}
