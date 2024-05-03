package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {

		fmt.Fprint(w, "ParseForm() err : %v", err)
		return

	}

	fmt.Fprintf(w, "POST Request succesful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func hellohandler(w http.ResponseWriter, r *http.Request) {

	// check if the url path and method are correct
	// return an err message if not
	if r.URL.Path != "/hello" {

		http.Error(w, "404 not found", http.StatusNotFound)
		return

	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hell0")

}

func main() {
	// create a fileserver that serves the static dir
	fileserver := http.FileServer(http.Dir("./static"))
	// register the fileserver
	http.Handle("./", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("starting server at port 8080\n")

	// start the server and listen for errors

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}

}
