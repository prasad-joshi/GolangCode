package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	program := os.Args[0]
	files := os.Args[1:]
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", program, err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			if _, ok := counts[line]; !ok {
				counts[line] = make(map[string]int)
			}
			counts[line][file]++
		}
	}

	for line, names := range counts {
		if len(names) == 1 {
			continue
		}

		fmt.Printf("Line: %s\n", line)
		for file, count := range names {
			fmt.Printf("\tFile:%s, Count=%d\n", file, count)
		}
	}
}
