package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/scraping/pkg/models"
	"sync"

	"github.com/gocolly/colly"
)

func Scrape(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Scraping")
}

func WriteToCSV(products []models.Product) {
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"URL", "Image", "Name", "Price"}
	writer.Write(headers)

	for _, product := range products {
		record := []string{
			product.Url,
			product.Image,
			product.Name,
			product.Price,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}

func NextPage(e *colly.HTMLElement, c *colly.Collector, visitedUrls *sync.Map) {
	nextPage := e.Attr("href")
	_, found := visitedUrls.Load(nextPage)
	if !found {
		visitedUrls.Store(nextPage, struct{}{})
		c.Visit(nextPage)
	}
}
