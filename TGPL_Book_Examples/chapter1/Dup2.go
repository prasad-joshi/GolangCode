package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	program := os.Args[0]
	files := os.Args[1:]
	if len(files) == 0 {
		CountLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			fd, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", program, err)
				continue
			}

			CountLines(fd, counts)
			fd.Close()
		}
	}

	for line, count := range counts {
		if count == 1 {
			continue
		}

		fmt.Printf("%d\t%s\n", count, line)
	}
}

func CountLines(fd *os.File, counts map[string]int) {
	input := bufio.NewScanner(fd)
	for input.Scan() {
		counts[input.Text()]++
	}
}