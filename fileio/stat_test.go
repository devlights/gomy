package fileio

import (
	"os"
	"testing"
)

func TestReadable(t *testing.T) {
	fi, err := os.Stat("testdata/0600.txt")
	if err != nil {
		t.Error(err)
	}

	if !Readable(fi) {
		t.Error("[want] true\t[got] false")
	}
}

func TestWritable(t *testing.T) {
	fi, err := os.Stat("testdata/0600.txt")
	if err != nil {
		t.Error(err)
	}

	if !Writable(fi) {
		t.Error("[want] true\t[got] false")
	}
}

func TestExecutable(t *testing.T) {
	fi, err := os.Stat("testdata/0700.txt")
	if err != nil {
		t.Error(err)
	}

	if !Executable(fi) {
		t.Error("[want] true\t[got] false")
	}
}

func TestRWX(t *testing.T) {
	fi, err := os.Stat("testdata/0700.txt")
	if err != nil {
		t.Error(err)
	}

	var (
		r = Readable(fi)
		w = Writable(fi)
		x = Executable(fi)
	)
	if !(r && w && x) {
		t.Errorf("[want] true,true,true\t[got] %v,%v,%v\n", r, w, x)
	}
}
