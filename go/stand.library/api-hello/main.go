// Go in action
// @jeffotoni

/// curl localhost:8080/api/hello
//////////////////////////////////////

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, welcome to the world, Go!"))
		}))
}
