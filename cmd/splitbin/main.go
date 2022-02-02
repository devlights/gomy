package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func run(r io.Reader, w io.Writer, length int) (err error) {
	for i := 1; ; i++ {
		var (
			buf   = make([]byte, length)
			count int
		)

		if count, err = r.Read(buf); err != nil {
			if err != io.EOF {
				err = fmt.Errorf("[ERR] Error while reading the record at %d(%v)\n", i, err)
				return
			}
			err = nil
			break
		}

		fmt.Fprintf(w, "%05d:%05d:%s\n", i, count, hex.EncodeToString(buf[:count]))
	}

	return
}

func open(path string) (rc io.ReadCloser, err error) {
	if path == "" {
		rc = io.NopCloser(os.Stdin)
		return
	}

	rc, err = os.Open(path)
	return
}

func args() (path string, length int, err error) {
	var (
		p = flag.String("f", "", "path to datafile (default stdin)")
		l = flag.Int("l", 16, "length to be split")
	)
	flag.Parse()

	path = *p
	length = *l

	if length < 1 {
		err = fmt.Errorf("invalid length value: %d (must be greater than 0)", length)
	}

	return
}

func main() {
	path, length, err := args()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rc, err := open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERR] Cannot open file: %s(%v)\n", path, err)
		os.Exit(2)
	}
	defer rc.Close()

	if err = run(rc, os.Stdout, length); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}
