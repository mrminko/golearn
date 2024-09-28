package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		args := r.URL
		fmt.Println(args)
		j, _ := json.Marshal("Hello")
		w.Write(j)
	})
	http.ListenAndServe("localhost:1111", nil)
}

//func outputWeight((w http.ResponseWriter, r *http.Request) int {
//	weight := 1000
//	return weight
//}
