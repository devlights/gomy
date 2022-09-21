package chans_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/devlights/gomy/chans"
)

func TestBufferContext(t *testing.T) {
	cases := []struct {
		name  string
		in    []interface{}
		count int
		out   [][]interface{}
	}{
		{"src[1,2,3]count[1]", []interface{}{1, 2, 3}, 1, [][]interface{}{{1}, {2}, {3}}},
		{"src[1,2,3]count[3]", []interface{}{1, 2, 3}, 3, [][]interface{}{{1, 2, 3}}},
		{"src[1,2,3]count[2]", []interface{}{1, 2, 3}, 2, [][]interface{}{{1, 2}, {3}}},
	}

	var (
		rootCtx = context.Background()
	)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				ctx, cxl = context.WithCancel(rootCtx)
			)
			defer cxl()

			results := make([][]interface{}, 0)
			for chunk := range chans.BufferContext(ctx, chans.Generator(ctx.Done(), c.in...), c.count) {
				results = append(results, chunk)
			}

			for i, chunk := range results {
				if !reflect.DeepEqual(chunk, c.out[i]) {
					t.Errorf("[want] %v\t[got] %v -- [index] %d", c.out[i], chunk, i)
				}
			}
		})
	}
}

func TestBuffer(t *testing.T) {
	cases := []struct {
		name  string
		in    []interface{}
		count int
		out   [][]interface{}
	}{
		{"src[1,2,3]count[1]", []interface{}{1, 2, 3}, 1, [][]interface{}{{1}, {2}, {3}}},
		{"src[1,2,3]count[3]", []interface{}{1, 2, 3}, 3, [][]interface{}{{1, 2, 3}}},
		{"src[1,2,3]count[2]", []interface{}{1, 2, 3}, 2, [][]interface{}{{1, 2}, {3}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			results := make([][]interface{}, 0)
			for chunk := range chans.Buffer(done, chans.Generator(done, c.in...), c.count) {
				results = append(results, chunk)
			}

			for i, chunk := range results {
				if !reflect.DeepEqual(chunk, c.out[i]) {
					t.Errorf("[want] %v\t[got] %v -- [index] %d", c.out[i], chunk, i)
				}
			}
		})
	}
}
