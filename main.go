package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", nil)
}
