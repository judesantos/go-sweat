package main

import (
	"fmt"

	"yourtechy.com/go-sweat/site_map/crawler"
)

func main() {

	c := crawler.NewCrawler()
	e := c.Crawl("https://yourtechy.com")

	if e != nil {
		panic(e)
	}

	fmt.Println("GetLinks return success!")
	fmt.Println()

	fmt.Println(c.ToXml())

}
