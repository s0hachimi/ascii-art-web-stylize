package main

import (
	"fmt"
	"net/http"

	assci "ascii-art-web/asciii"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", assci.Index)
	http.HandleFunc("/ascii-art", assci.Result)
	fmt.Println("Server is listening on port 8080...   http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
