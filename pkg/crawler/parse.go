package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Parse(url string) {

	fmt.Println("parsing... ", url)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("developer.mozilla.org"),
	)

	// var wg sync.WaitGroup

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

}
