package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	program := os.Args[0]
	files := os.Args[1:]
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", program, err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count == 1 {
			continue
		}

		fmt.Printf("%d\t%s\n", count, line)
	}
}
