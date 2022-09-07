package chans

import "context"

// SliceContext は、Slice の context.Context 版です.
func SliceContext[T any](ctx context.Context, in <-chan T) []T {
	return Slice(ctx.Done(), in)
}

// Slice は、指定されたチャネルから情報を読み取りスライスにして返します.
//
// 本処理は、done もしくは in のどちらかがクローズされるまで結果を返しません。
func Slice[T any](done <-chan struct{}, in <-chan T) []T {
	var (
		ret = make([]T, 0)
	)

LOOP:
	for {
		select {
		case <-done:
			break LOOP
		case v, ok := <-in:
			if !ok {
				break LOOP
			}
			ret = append(ret, v)
		}
	}

	return ret
}
