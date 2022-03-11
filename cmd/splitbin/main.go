package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func run(r io.Reader, w io.Writer, length int) error {
	for i := 1; ; i++ {
		buf := make([]byte, length)
		
		count, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("[ERR] Error while reading the record at %d(%v)", i, err)
			}
			break
		}

		fmt.Fprintf(w, "%05d:%05d:%s\n", i, count, hex.EncodeToString(buf[:count]))
	}

	return nil
}

func args() (length int, err error) {
	var (
		l = flag.Int("l", 16, "length to be split")
	)
	flag.Parse()

	length = *l
	if length < 1 {
		err = fmt.Errorf("invalid length value: %d (must be greater than 0)", length)
	}

	return
}

func main() {
	length, err := args()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = run(os.Stdin, os.Stdout, length)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(9)
	}
}
