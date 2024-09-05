package main

import (
	"fmt"
	"net/http"
)

var points int

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/click", clickHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func clickHandler(w http.ResponseWriter, r *http.Request) {
	points++
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Points: %d", points)
}
