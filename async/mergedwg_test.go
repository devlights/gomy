package async

import (
	"sync"
	"testing"
	"time"
)

func TestMergedWaitGroup(t *testing.T) {
	type (
		testin struct {
			wgCount   int
			waitTimes []time.Duration
		}
		testout struct {
			estimation time.Duration
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				wgCount: 1,
				waitTimes: []time.Duration{
					100 * time.Millisecond,
				},
			},
			out: testout{estimation: 110 * time.Millisecond},
		},
		{
			in: testin{
				wgCount: 3,
				waitTimes: []time.Duration{
					100 * time.Millisecond,
					200 * time.Millisecond,
					300 * time.Millisecond,
				},
			},
			out: testout{estimation: 310 * time.Millisecond},
		},
		{
			in: testin{
				wgCount: 5,
				waitTimes: []time.Duration{
					100 * time.Millisecond,
					200 * time.Millisecond,
					300 * time.Millisecond,
					400 * time.Millisecond,
					500 * time.Millisecond,
				},
			},
			out: testout{estimation: 510 * time.Millisecond},
		},
	}

	for caseIndex, c := range cases {
		func() {
			wgList := make([]*sync.WaitGroup, 0)

			for i := 0; i < c.in.wgCount; i++ {
				wg := sync.WaitGroup{}
				wgList = append(wgList, &wg)

				wg.Add(1)
				go func(wg *sync.WaitGroup, waitTime time.Duration, index int) {
					defer wg.Done()

					defer t.Logf("[wg-%02d] wait done : %v", index, waitTime)
					t.Logf("[wg-%02d] wait start: %v", index, waitTime)

					<-time.After(waitTime)
				}(&wg, c.in.waitTimes[i], i)
			}

			mwg := NewMergedWaitGroup(wgList...)

			start := time.Now()
			mwg.Wait()
			elapsed := time.Since(start)

			t.Logf("[test-%02d][wgCount=%d]\telapsed: %v\testimation: %v", caseIndex, c.in.wgCount, elapsed, c.out.estimation)
			if c.out.estimation < elapsed {
				t.Errorf("want: <- %v\tgot: %v", c.out.estimation, elapsed)
			}
		}()
	}
}
