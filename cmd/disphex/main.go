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

func args() (sep string, col int, err error) {
	var (
		s = flag.String("sep", ":", "separator")
		c = flag.Int("col", 2, "column number (zero start)")
	)

	flag.Parse()

	sep = *s
	col = *c

	if col < 0 {
		err = fmt.Errorf("invalid col number: %d", col)
		return
	}

	return
}

func main() {
	sep, col, err := args()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = run(os.Stdin, os.Stdout, sep, col); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(9)
	}
}
