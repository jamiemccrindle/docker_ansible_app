package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello London Docker Meetup!")
		})
	http.ListenAndServe(":8000", nil)
}
