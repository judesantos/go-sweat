package main

import (
	"fmt"

	"github.com/judesantos/go-bookstore_utils/logger"
	"yourtechy.com/go-sweat/link_parser/links"
)

func main() {

	fmt.Printf("Start link parser\n\n")

	fmt.Printf("Test 1: Get Hello Links\n")
	GetLinks("./tests/data/hello.html")

	fmt.Printf("\nTest 2: Get Adjacent Links\n")
	GetLinks("./tests/data/adjacent-links.html")

	fmt.Printf("\nTest 3: Get Sectioned Links\n")
	GetLinks("./tests/data/sectioned-links.html")

	fmt.Printf("\nTest 4: Get Nested Links\n")
	GetLinks("./tests/data/nested-links.html")
}

func GetLinks(source string) {

	var _links *[]*links.Link
	var err error

	_links, err = links.GetLinks(source)

	if err != nil {
		logger.Error("GetLinks failed:", err)
	}

	for i, link := range *_links {
		fmt.Printf("  %d.) URL: %s, Text: %s\n", i+1, link.Url, link.Text)
	}

}
