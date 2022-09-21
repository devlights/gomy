package chans

import "context"

// MergeContext は、 Merge の context.Context 版です。
func MergeContext[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return FanInContext(ctx, channels...)
}

// Merge は、複数のチャネルシーケンスを纏めて一つのチャネルから出力するようにします。
// 取り出される順番は不定です。FanIn関数のエイリアスです。
//
// 出力順を守りたい場合は Concat を利用してください。
func Merge[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	return FanIn(done, channels...)
}
