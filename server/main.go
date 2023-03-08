package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS!")
	})

	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}