package chans

// WhenAny -- 指定した１つ以上のチャネルのどれかが１つが閉じられたら、閉じるチャネルを返します。
//
// チャネルを一つも渡さずに呼び出すと、既に close 済みのチャネルを返します。
func WhenAny[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		nilCh := make(chan T)
		close(nilCh)

		return nilCh
	case 1:
		return channels[0]
	}

	orDone := make(chan T)
	go func() {
		defer close(orDone)

		// 再帰呼出しの回数を抑えるために len が (2 or 3) のときは再帰せずに済ませる
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		case 3:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-WhenAny(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}
