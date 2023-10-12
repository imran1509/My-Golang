package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to Mod in Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("Hello mod users")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang series on youtube</h1>"))
}
