package chans

import "context"

// GeneratorContext は、Generator の context.Context 版です.
func GeneratorContext[T any](ctx context.Context, in ...T) <-chan T {
	return Generator(ctx.Done(), in...)
}

// Generator -- 指定されたデータを出力するチャネルを生成します。
//
// ForEach関数のエイリアスです。
func Generator[T any](done <-chan struct{}, in ...T) <-chan T {
	return ForEach(done, in...)
}
