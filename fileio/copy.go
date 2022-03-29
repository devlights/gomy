package fileio

import (
	"bufio"
	"io"
	"os"
)

// Copy -- ファイルをコピーします。
//
// REFERENCES
//  - https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file
func Copy(src, dst string) (err error) {
	var (
		in  *os.File
		out *os.File
	)

	if in, err = os.Open(src); err != nil {
		return err
	}
	defer func() { err = in.Close() }()

	if out, err = os.Create(dst); err != nil {
		return err
	}
	defer func() { err = out.Close() }()

	var (
		r = bufio.NewReader(in)
		w = bufio.NewWriter(out)
	)

	if _, err = io.Copy(w, r); err != nil {
		return err
	}

	if err = w.Flush(); err != nil {
		return err
	}

	return nil
}
