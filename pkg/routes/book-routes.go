package routes

/*
This file is used for routing when the client hits the api
*/

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/GO-Book-Management-System/pkg/controllers"
)

var RegisterBookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/book/all", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}/update", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}/delete", controllers.DeleteBook).Methods("DELETE")
}
