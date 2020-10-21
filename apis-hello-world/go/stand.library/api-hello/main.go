// Go in action
// @jeffotoni

/// curl -i localhost:8080/hello
//////////////////////////////////////

package main

import (
	"net/http"
)

func main() {
	println("Starting server at 0.0.0.0:8080")
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, welcome to the world, Go!"))
		}))
}
