///////build a simple server/////////////
package main

import (
	. "fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "The method is not supported", http.StatusNotFound)
		return
	}
	Fprint(w, "hello programmer!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	Fprint(w, "POST request is successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	Fprintf(w, "Name=%s\n", name)
	Fprintf(w, "Address=%s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	Printf("Starting server at port 9090\n")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
