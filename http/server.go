package http

import (
	"fmt"
	"log"
	"net/http"
)

// Start the server
func Start() {
	// http.HandleFunc("/", hello)
	// http.Handler("/api", routes())

	log.Fatal(http.ListenAndServe(":8080", routes()))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
