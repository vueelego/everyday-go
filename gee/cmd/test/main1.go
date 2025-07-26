package main

import (
	"fmt"
	"net/http"
)

func main1() {
	http.HandleFunc("/", indexHandler1)
	http.HandleFunc("/hello", helloHandler1)
	http.ListenAndServe(":4000", nil)
}

func indexHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, fmt.Sprintf("Header[%q] = %q\n", k, v))
	}
}
