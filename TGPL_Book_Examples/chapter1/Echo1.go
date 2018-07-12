package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for range [10000]int{} {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}
	fmt.Printf("Time %.2fs\n", time.Since(start).Seconds())
}
