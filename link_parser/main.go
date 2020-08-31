package main

import (
	"fmt"
	"io"
	"os"

	"yourtechy.com/go-sweat/link_parser/links"
	"yourtechy.com/go-sweat/utils/logger"
)

var (
	log = logger.NewLogger()
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

func getReader(source string) io.Reader {

	r, err := os.Open(source)

	if err != nil {
		fmt.Println("os.Open failed!", err)
		return nil
	}

	return io.Reader(r)
}

func GetLinks(source string) {

	var _links *[]*links.Link
	var err error

	r := getReader(source)
	if r == nil {
		panic(fmt.Sprintf("getReader failed to load file %s", source))
	}

	_links, err = links.GetLinks(&r)

	if err != nil {
		log.Error("GetLinks failed:", err)
	}

	fmt.Println("links found:", len(*_links))

	for i, link := range *_links {
		fmt.Printf("  %d.) URL: %s, Text: %s\n", i+1, link.Url, link.Text)
	}

}
