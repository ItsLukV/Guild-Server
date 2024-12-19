package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Haha get ratted")
	})

	fmt.Println("Serving on port 5555")
	http.ListenAndServe(":5555", nil)
}
