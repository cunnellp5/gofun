package main

import (
	"fmt"
	"project/scraping/pkg/controllers"
	"project/scraping/pkg/models"
	"project/scraping/pkg/utils"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	products := []models.Product{}
	visitedUrls := sync.Map{}

	c := utils.SetupColly()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := models.Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img", "src"),
			Name:  e.ChildText("h2"),
			Price: e.ChildText("span.price"),
		}
		products = append(products, product)
	})

	c.OnHTML("a.next", func(e *colly.HTMLElement) {
		controllers.NextPage(e, c, &visitedUrls)
	})

	// Set up the OnScraped callback with access to the products slice
	c.OnScraped(func(r *colly.Response) {
		controllers.WriteToCSV(products)
	})

	// Now visit the URL
	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
