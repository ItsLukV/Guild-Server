package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	fmt.Println("Serving on port 5555")
	http.ListenAndServe(":5555", nil)
}
