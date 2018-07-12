package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for range [10000]int{} {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
	fmt.Printf("Time %.2fs\n", time.Since(start).Seconds())
}
