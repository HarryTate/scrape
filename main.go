package main

import (
	"fmt"
	"log"
	"os"

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
		heading := e.ChildText(".hero-home__heading")
		f, err := os.Create("data.txt")

		if err != nil {
			log.Fatal(err)
		}
		_, err2 := f.WriteString(heading)

		defer f.Close()

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println(heading)
	})

	startUrl := fmt.Sprintf("https://www.harrytate.co.uk")
	c.Visit(startUrl)
}
