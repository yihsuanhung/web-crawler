package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Parse(url string) string {

	fmt.Println("parsing... ", url)

	// Instantiate default collector
	c := colly.NewCollector()

	var data string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		data += e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	return data
}
