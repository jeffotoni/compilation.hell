package main

import "net/http"

func main() {
	println("Starting server at 0.0.0.0:8080")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, welcome to the world, Go!"))
	})
	http.ListenAndServe(":8080", nil)
}
