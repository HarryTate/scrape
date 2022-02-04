package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	crawl()
}

func crawl() {
	c := colly.NewCollector(
		colly.AllowedDomains("harrytate.co.uk", "www.harrytate.co.uk"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		heading := e.ChildText(".skill_title")
		fmt.Println(heading)
	})

	startUrl := fmt.Sprintf("https://www.harrytate.co.uk")
	c.Visit(startUrl)
}
