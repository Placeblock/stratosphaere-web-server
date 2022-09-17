package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hellp World")
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":9130", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Schue, Hose, Hemd"))
}
