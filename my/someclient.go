package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", client)
	http.ListenAndServe("localhost:2222", nil)
}

func client(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("localhost:1111")
	fmt.Println(res)
	io.WriteString(w, "hello form client")
}
