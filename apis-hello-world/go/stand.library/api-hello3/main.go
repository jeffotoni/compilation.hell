package main

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, welcome to the world, Go!"))
}

func main() {
	println("Starting server at 0.0.0.0:8080")
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(Hello))
	server :=
		&http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
	server.ListenAndServe()
}
