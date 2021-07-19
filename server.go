package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func randomnum(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello This is your random number: %v\n", rand.Float64())
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/random", randomnum)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8882", nil)
}
