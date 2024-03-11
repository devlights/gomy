package promise

type (
	// Future[T] は、「今はまだ得られていないが将来得られるはずの入力」を表します。
	Future[T any] struct {
		value T
		wait  chan struct{}
	}
)

// IsDone は、この Future[T] が完了したかどうかを返します。
func (me *Future[T]) IsDone() bool {
	select {
	case <-me.wait:
		return true
	default:
		return false
	}
}

func (me *Future[T]) set(v T) {
	if me.IsDone() {
		return
	}

	me.value = v
	close(me.wait)
}

// Get は、この Future[T] の結果を返します。
func (me *Future[T]) Get() T {
	<-me.wait
	return me.value
}
