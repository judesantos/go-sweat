package links

import (
	"errors"
	"fmt"

	"github.com/judesantos/go-bookstore_utils/logger"
	parser "yourtechy.com/go-sweat/link_parser/parser/html"
)

// Link - The link container the link URL and string content
type Link struct {
	Url  string
	Text string
}

// GetLinks - Get all hyperlinks from a given html file
func GetLinks(source string) (*[]*Link, error) {

	parser, err := parser.NewHtmlParser(source)

	if err != nil {
		logger.Error("Get parser failed!", err)
		return nil, err
	}

	node := parser.FindNode("body")

	if node == nil {
		fmt.Println("Find body failed")
		return nil, errors.New("Find body failed")
	}

	links := []*Link{}

	buildLinks(node, &links)

	return &links, nil
}

// buildLinks - Iterate through each node, children, and siblings,
//              check if href is found, extract and store in an
//              array of links.
//
// @params: node - The current node
//          links - Returns all found links.
func buildLinks(node *parser.HtmlNode, links *[]*Link) {

	if node.GetNodeName() == "a" {

		attrib := node.GetAttribute("href")
		text := node.GetText(0)

		*links = append(*links, &Link{
			Url:  attrib.Value,
			Text: text,
		})

	}

	var child *parser.HtmlNode

	if child = node.GetChild(); child != nil {
		buildLinks(child, links)
	}

	if child = node.NextSibling(); child != nil {
		buildLinks(child, links)
	}
}
