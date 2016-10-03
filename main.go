package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":80", nil))
}
