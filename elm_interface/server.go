package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Serving files at port: \"%v\"\n", os.Args[1])
	http.ListenAndServe(os.Args[1], nil)
}
