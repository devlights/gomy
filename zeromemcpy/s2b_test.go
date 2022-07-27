package zeromemcpy_test

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/devlights/gomy/zeromemcpy"
)

func TestS2B(t *testing.T) {
	type tdata struct {
		in  string
		out []byte
	}

	testdata := []tdata{
		{"", []byte("")},
		{"hello", []byte("hello")},
		{"hello     world", []byte("hello     world")},
		{"こんにちわ世界", []byte("こんにちわ世界")},
	}

	for i, data := range testdata {
		i := i
		data := data

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			in := data.in
			expected := data.out

			result := zeromemcpy.S2B(in)
			if !bytes.Equal(result, expected) {
				t.Errorf("[want] %v\t[got] %v", expected, result)
			}
		})
	}
}

func BenchmarkS2BNormalWay(b *testing.B) {
	var (
		buf strings.Builder
	)

	for i := 0; i < 1_000_000; i++ {
		buf.WriteString(strconv.Itoa(i))
	}

	var (
		target = buf.String()
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.Discard.Write([]byte(target))
	}
}

func BenchmarkS2B(b *testing.B) {
	var (
		buf strings.Builder
	)

	for i := 0; i < 1_000_000; i++ {
		buf.WriteString(strconv.Itoa(i))
	}

	var (
		target = buf.String()
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.Discard.Write(zeromemcpy.S2B(target))
	}
}
