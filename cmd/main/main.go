package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/GO-Book-Management-System/pkg/routes"
)

const url = "localhost:8000"

func main() {
	r := mux.NewRouter()               // create a new router using gorilla mux
	routes.RegisterBookStoreRouters(r) // pass the new router to the routes we created
	http.Handle("/", r)
	fmt.Printf("Startng server at %v", url)
	log.Fatal(http.ListenAndServe(url, r)) // run the server. log.fatel will throw error
}
