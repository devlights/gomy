package iter_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/devlights/gomy/iter"
)

func ExampleRange() {
	var (
		appLog = log.New(os.Stdout, "", 0)
	)

	for i := range iter.Range(3) {
		appLog.Print(i)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleRangeFn() {
	var (
		appLog = log.New(os.Stdout, "", 0)
	)

	fn := func(i int) error {
		appLog.Print(i)
		return nil
	}

	_, err := iter.RangeFn(3, fn)
	if err != nil {
		appLog.Fatal(err)
	}

	// Output:
	// 0
	// 1
	// 2
}

func TestRange(t *testing.T) {
	cases := []struct {
		name string
		in   int
		out  []int
	}{
		{"loop-3", 3, []int{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var ans []int
			for i := range iter.Range(c.in) {
				ans = append(ans, i)
			}

			r1, r2 := toStringSlice(c.out, ","), toStringSlice(ans, ",")
			if r1 != r2 {
				t.Errorf("[want] %s\t[got] %s", r1, r2)
			}
		})
	}
}

func TestRangeFn(t *testing.T) {
	var (
		testErr = fmt.Errorf("test error")
	)

	cases := []struct {
		name string
		in   int
		outi int
		oute error
	}{
		{"loop-3", 3, 0, testErr},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ri, err := iter.RangeFn(c.in, func(i int) error {
				t.Log(i)
				return testErr
			})

			if ri != c.outi {
				t.Errorf("[want] %v\t[got] %v", c.outi, ri)
			}

			if err != c.oute {
				t.Errorf("[want] %v\t[got] %v", c.oute, err)
			}
		})

	}
}

func toStringSlice(ints []int, delimiter string) string {
	var strs []string
	for v := range ints {
		strs = append(strs, strconv.Itoa(v))
	}

	return strings.Join(strs, delimiter)
}
