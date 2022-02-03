package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		b = flag.Int("b", 0, "begin (starts from 0)")
		c = flag.Int("c", 1, "count")
	)
	flag.Parse()

	var (
		scanner = bufio.NewScanner(os.Stdin)
		begin   = *b
		end     = (begin + (*c)) - 1
	)
	for i := 0; scanner.Scan(); i++ {
		if begin <= i && i <= end {
			txt := scanner.Text()
			fmt.Fprintln(os.Stdout, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
