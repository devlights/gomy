package ctxs

import "context"

// ToDoneCh は、指定された context.Context から done チャネルを取り出してスライスで返します.
//
// 取得した done チャネルのスライスを chans.WhenAll() に指定すると全コンテキストが完了するまで
// 待機することが出来ます。
//
// Example:
//
//	<-chans.WhenAll(ctxs.ToDoneCh(contexts...)...) // contexts is []context.Context
func ToDoneCh(contexts ...context.Context) []<-chan struct{} {
	dones := make([]<-chan struct{}, len(contexts))

	for i, ctx := range contexts {
		dones[i] = ctx.Done()
	}

	return dones
}
