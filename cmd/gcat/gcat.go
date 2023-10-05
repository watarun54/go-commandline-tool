package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var nflag bool

func cat(file *os.File, n *int) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		l := s.Text()
		if nflag {
			fmt.Printf("%d  %s\n", *n, l)
			*n++
		} else {
			fmt.Println(l)
		}
	}
}

func main() {
	flag.BoolVar(&nflag, "n", false, "行番号を表示")
	flag.Parse()

	n := 1
	for _, name := range flag.Args() {
		file, err := os.Open(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cat(file, &n)

		file.Close()
	}
}
