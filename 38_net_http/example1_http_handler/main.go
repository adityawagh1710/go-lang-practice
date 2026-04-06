package main

import (
	"fmt"
	"net/http"
)

func MyHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func MyHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello John\n")
}

func main() {

	//  sudo lsof -i :8080
	//  sudo ss -tlnp | grep 8080
	//  sudo kill -9 <pid>
	//  curl http://localhost:8080/
	//  curl http://localhost:8080/John

	http.HandleFunc("/", MyHandler1)

	http.HandleFunc("/John", MyHandler2)

	http.ListenAndServe(":8080", nil)
}
