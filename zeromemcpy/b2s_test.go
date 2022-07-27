package zeromemcpy_test

import (
	"bytes"
	"io"
	"strconv"
	"testing"

	"github.com/devlights/gomy/zeromemcpy"
)

func TestB2S(t *testing.T) {
	type tdata struct {
		in  []byte
		out string
	}

	testdata := []tdata{
		{[]byte(""), ""},
		{[]byte("hello"), "hello"},
		{[]byte("hello     world"), "hello     world"},
		{[]byte("こんにちわ世界"), "こんにちわ世界"},
	}

	for i, data := range testdata {
		i := i
		data := data

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			in := data.in
			expected := data.out

			result := zeromemcpy.B2S(in)
			if result != expected {
				t.Errorf("[want] %v\t[got] %v", expected, result)
			}
		})
	}
}

func BenchmarkB2SNormalWay(b *testing.B) {
	var (
		buf = new(bytes.Buffer)
	)

	for i := 0; i < 1_000_000; i++ {
		buf.WriteString(strconv.Itoa(i))
	}

	var (
		target = buf.Bytes()
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.WriteString(io.Discard, string(target))
	}
}

func BenchmarkB2S(b *testing.B) {
	var (
		buf = new(bytes.Buffer)
	)

	for i := 0; i < 1_000_000; i++ {
		buf.WriteString(strconv.Itoa(i))
	}

	var (
		target = buf.Bytes()
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.WriteString(io.Discard, zeromemcpy.B2S(target))
	}
}
