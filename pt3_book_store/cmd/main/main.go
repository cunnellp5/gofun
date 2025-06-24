package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cunnellp5/pt3_book_store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("fugging listening bro 🌽")
	log.Fatal(http.ListenAndServe(":9010", r))
}
