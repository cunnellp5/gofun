package utils

import (
	"log"

	"github.com/gocolly/colly"
)

func SetupColly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	// set up the user agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	// set up the proxy
	err := c.SetProxy("http://51.81.245.3:17981")
	if err != nil {
		log.Fatal(err)
	}

	return c
}
