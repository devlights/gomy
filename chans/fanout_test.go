package chans

import (
	"sort"
	"testing"
	"time"

	"github.com/devlights/gomy/output"
)

func TestFanOut(t *testing.T) {
	type (
		result struct {
			index int
			value interface{}
		}
	)

	var (
		givenTime    = 1 * time.Second
		numGoroutine = 2
		items        = []interface{}{"hello", "world", "こんにちわ", "世界"}
		results      = make([]*result, 0, 0)
	)
	var (
		done  = make(chan struct{})
		outCh = make(chan interface{})
	)

	defer close(done)

	// 処理するのに t に指定された時間がかかる関数
	fn := func(item interface{}, t time.Duration) {
		<-time.After(t)
		output.Stdoutl("[処理]", item)
	}

	// パイプライン生成
	forEachCh := ForEach(done, items...)
	enumerateCh := Enumerate(done, forEachCh)
	doneChList := FanOut(done, enumerateCh, numGoroutine, func(e interface{}) {
		if v, ok := e.(*IterValue); ok {
			fn(v.Value, givenTime)
			outCh <- &result{
				index: v.Index,
				value: v.Value,
			}
		}
	})

	// 処理完了とともに出力用チャネルを閉じる
	go func() {
		defer close(outCh)
		<-WhenAll(doneChList...)
	}()

	// 結果を吸い出し
	for v := range outCh {
		results = append(results, v.(*result))
	}

	// 正しい順序に並び替え
	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	for i := 0; i < len(items); i++ {
		l := items[i]
		r := results[i].value

		if l != r {
			t.Errorf("want: %v\tgot: %v", l, r)
		}
	}
}
