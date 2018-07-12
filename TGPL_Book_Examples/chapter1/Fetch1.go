package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	prog := os.Args[0]

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"%s: fetching url '%s' failed with error '%v'\n",
				prog, url, err)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"%s: reading url '%s' body failed with error '%v'\n",
				prog, url, err)
			continue
		}

		fmt.Printf("%s", b)
	}
}
