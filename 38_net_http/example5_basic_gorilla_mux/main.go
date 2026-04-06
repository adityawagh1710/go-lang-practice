package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/teststts/{id}", handler)

	http.ListenAndServe(":8080", r)
}
