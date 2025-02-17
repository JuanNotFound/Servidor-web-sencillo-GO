package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Start server at port 8000\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprint(w, "POST request succes!")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Welcome %s! \n you live in: %s", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	_, err := fmt.Fprintf(w, "Welcome to simple Go Server!")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error on our side!", http.StatusInternalServerError)
		return
	}

}
