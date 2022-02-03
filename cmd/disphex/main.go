package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func run(r io.Reader, w io.Writer, sep string, col int) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var (
			txt   = scanner.Text()
			parts = strings.Split(txt, sep)
		)

		if len(parts)-1 < col {
			fmt.Fprintln(w, "")
			continue
		}

		var (
			data         = parts[col]
			decoded, err = hex.DecodeString(data)
		)

		if err != nil {
			return fmt.Errorf("error at decoding process (%s)", data)
		}

		for _, b := range decoded {
			fmt.Fprintf(w, "%02x ", b)
		}
		fmt.Fprintln(w, "")
	}

	return scanner.Err()
}

func open(path string) (rc io.ReadCloser, err error) {
	if path == "" {
		rc = io.NopCloser(os.Stdin)
		return
	}

	rc, err = os.Open(path)
	return
}

func args() (path string, sep string, col int, err error) {
	var (
		p = flag.String("f", "", "path to datafile (default stdin)")
		s = flag.String("sep", ":", "separator")
		c = flag.Int("col", 2, "column number (zero start)")
	)

	flag.Parse()

	path = *p
	sep = *s
	col = *c

	if col < 0 {
		err = fmt.Errorf("invalid col number: %d", col)
		return
	}

	return
}

func main() {
	path, sep, col, err := args()
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

	if err = run(rc, os.Stdout, sep, col); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}
