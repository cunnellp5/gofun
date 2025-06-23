package routes

import (
	"github.com/cunnellp5/pt3_book_store/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Method("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Method("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Method("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Method("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Method("Delete")
}
