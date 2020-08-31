/*
	Crawler - Provided a base URL, crawler will read each site page and extract all
	relevant URLs, excluding external links.

	sourceFile crawler.go
*/
package crawler

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"strings"

	"yourtechy.com/go-sweat/link_parser/links"
	"yourtechy.com/go-sweat/utils/logger"
)

var (
	log = logger.NewLogger()
)

type crawler struct {
	baseUrl   string
	links     []string
	processed map[string]struct{}
}

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

/************************************************************************
 *
 * Public types, methods
 *
 ***********************************************************************/

// NewCrawler - Create new crawler instance
func NewCrawler() *crawler {
	return &crawler{
		baseUrl:   "",
		links:     make([]string, 0),
		processed: map[string]struct{}{},
	}
}

// Crawl - Get all links from the target site
func (c *crawler) Crawl(baseUrl string) error {
	err := getLinks(c, baseUrl)
	if err != nil {
		return err
	}

	return nil
}

// ToArray - Return links in a string array
//				   Call after a sucessful Crawl() response
func (c *crawler) ToArray() *[]string {
	return &c.links
}

// ToXml - Return string in sitemap protocol xml format
//				 Call after a sucessful Crawl() response
func (c *crawler) ToXml() (string, error) {

	const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

	toXml := urlset{
		Xmlns: xmlns,
	}

	for _, link := range c.links {
		toXml.Urls = append(toXml.Urls, loc{link})
	}

	var strBuf bytes.Buffer
	enc := xml.NewEncoder(&strBuf)
	enc.Indent("", " ")

	if err := enc.Encode(toXml); err != nil {
		return "", err
	}

	return strBuf.String(), nil
}

/************************************************************************
 *
 * Private methods
 *
 ***********************************************************************/

// GetLinks - Visit all pages from the specified siteUrl and collect all site
//            links. Exclude external links
func getLinks(c *crawler, siteUrl string) error {

	var _links *[]*links.Link

	{
		// fetch
		res, err := http.Get(siteUrl)

		if err != nil {
			return err
		}

		defer res.Body.Close()

		// mark found link as processed

		if c.baseUrl == "" {
			// Save actual base URL
			reqUrl := res.Request.URL
			baseUrl := &url.URL{
				Scheme: reqUrl.Scheme,
				Host:   reqUrl.Host,
			}
			c.baseUrl = baseUrl.String()
			c.processed[c.baseUrl] = struct{}{}
		} else {
			c.processed[siteUrl] = struct{}{}
		}

		r := io.Reader(res.Body)
		_links, err = links.GetLinks(&r)

		if err != nil {
			return err
		}
	}

	// process each page and get more links

	for _, link := range *_links {

		// If relative path, we need to prefix baseUrl
		if strings.HasPrefix(link.Url, "/") {
			link.Url = c.baseUrl + link.Url
		}

		if _, exists := c.processed[link.Url]; exists {
			// Ignore if we already have a record of this link
			continue
		}

		// Ignore invalid links (e.g: #, javascript(void))
		if 0 == strings.Index(link.Url, c.baseUrl) {

			// Valid internal url - record link
			c.links = append(c.links, link.Url)
			// Fetch more links - extracet links from this sub-link/page
			getLinks(c, link.Url)
		}

	} // end for

	return nil
}
