package chans

import "context"

// ConcatContext は、 Concat の context.Context 版です.
func ConcatContext[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return Concat(ctx.Done(), channels...)
}

// Concat は、指定されたチャネルのシーケンスを順に消費していく単一のチャネルを返します.
func Concat[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	if len(channels) == 0 {
		c := make(chan T)
		close(c)

		return c
	}

	sequences := make(chan (<-chan T), len(channels))
	func() {
		defer close(sequences)

		for _, c := range channels {
			sequences <- c
		}
	}()

	return Bridge(done, sequences)
}
