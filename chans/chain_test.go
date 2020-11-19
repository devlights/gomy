package chans

import (
	"strings"
	"testing"
	"time"
)

func TestChain(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		out  []string
	}{
		{"2 groutines chain", []string{"hello", "world"}, []string{"hello", "WORLD"}},
		{"3 groutines chain", []string{"hello", "world", "one"}, []string{"hello", "WORLD", "one"}},
		{"4 groutines chain", []string{"hello", "world", "one", "two"}, []string{"hello", "WORLD", "one", "TWO"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				done   = make(chan struct{})
				result = make(chan string)
				ch     <-chan struct{}
			)

			defer close(done)

			for i, v := range c.in {
				v := v

				if ch == nil {
					// 1st goroutine
					ch = func(v string) <-chan struct{} {
						done := make(chan struct{})
						go func() {
							defer close(done)
							result <- v
						}()
						return done
					}(v)

					continue
				}

				// chain 2nd, 3rd...more
				index := i + 1
				ch = Chain(done, ch, func(ft time.Time) {
					if index%2 == 0 {
						v = strings.ToUpper(v)
					}

					t.Logf("[goroutine][%v]\t%v\t%v", index, v, ft.UTC().Unix())
					result <- v
				})
			}

			go func() {
				defer close(result)
				<-ch
				t.Log("[chain goroutines] done")
			}()

			var i int
			for v := range result {
				if v != c.out[i] {
					t.Errorf("[want] %v\t[got] %v", c.out[i], v)
				}

				i++
			}
		})
	}
}
