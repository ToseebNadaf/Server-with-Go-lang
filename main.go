package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	fname := r.FormValue("fname")
	lastName := r.FormValue("lname")
	fmt.Fprintf(w, "Name: %s\n", fname)
	fmt.Fprintf(w, "Last Name: %s\n", lastName)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not valid", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello !")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is starting on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}