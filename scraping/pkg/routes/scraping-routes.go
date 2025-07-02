package routes

import (
	"project/scraping/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterScrapingRoutes = func(router *mux.Router) {
	router.HandleFunc("/scrape", controllers.Scrape).Methods("GET")
}
