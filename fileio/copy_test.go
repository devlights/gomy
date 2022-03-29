package fileio

import (
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	src := "gomy-testcopy-src"
	dst := "gomy-testcopy-dst"

	cleanup := func(ignore bool) {
		for _, f := range []string{src, dst} {
			if err := os.Remove(f); err != nil {
				if !ignore {
					t.Errorf("[want] nil\t[got] %v", err)
				}
			}
		}
	}

	cleanup(true)
	defer cleanup(false)

	f, err := os.Create(src)
	if err != nil {
		t.Errorf("[want] nil\t[got] %v", err)
	}

	f.WriteString("hello world")
	f.Close()

	if err := Copy(src, dst); err != nil {
		t.Errorf("[want] nil\t[got] %v", err)
	}
}
