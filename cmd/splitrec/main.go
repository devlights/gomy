package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func toS(items []int) string {
	var s string

	for _, item := range items {
		s = fmt.Sprintf("%s%d,", s, item)
	}

	return s[:len(s)-1]
}

func run(r io.Reader, w io.Writer, sep string, col int, fields []int) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var (
			txt   = scanner.Text()
			parts = strings.Split(txt, sep)
		)

		if len(parts)-1 < col {
			return fmt.Errorf("there are not enough columns for processing (%d)", len(parts))
		}

		var (
			data         = parts[col]
			decoded, err = hex.DecodeString(data)
		)

		if err != nil {
			return fmt.Errorf("error at decoding process (%s)", data)
		}

		if len(fields) == 0 {
			fmt.Fprintln(w, data)
			continue
		}

		fmt.Fprintf(w, "%s:%s:%s:", parts[0], parts[1], toS(fields))

		var curr int
		for _, field := range fields {
			next := curr + field
			if len(decoded) < next {
				break
			}
			fmt.Fprintf(w, "%02x:", decoded[curr:next])
			curr = next
		}
		fmt.Fprintln(w, "")
	}

	return scanner.Err()
}

func args() (sep string, col int, fields []int, err error) {
	var (
		f = flag.String("fields", "", "field-list (ex: 20,2,2,4)")
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

	fs := strings.Split(*f, ",")
	for _, v := range fs {
		i, e := strconv.Atoi(v)
		if e != nil {
			err = e
			return
		}
		fields = append(fields, i)
	}

	return
}

func main() {
	sep, col, fields, err := args()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = run(os.Stdin, os.Stdout, sep, col, fields); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}
