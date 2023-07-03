package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusFound)
	}

	fmt.Fprint(w, "Hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Adress = %s", adress)

}

func main() {
	fileSever := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileSever)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}

}
