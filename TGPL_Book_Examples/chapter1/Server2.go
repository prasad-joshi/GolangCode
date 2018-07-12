// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var counter int

func main() {
	http.HandleFunc("/", EchoHandler)
	http.HandleFunc("/counter", CounterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func CounterHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	c := counter
	mutex.Unlock()
	fmt.Fprintf(writer, "Count %d\n", c)
}
