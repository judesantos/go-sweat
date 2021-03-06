package main

import (
	"fmt"
	"time"

	"yourtechy.com/go-sweat/site_map/crawler"
)

func main() {

	var isDone bool = false
	var url = "https://yourtechy.com"

	// create crawler instance
	c := crawler.NewCrawler()

	// run progress counter in background
	go func() {
		fmt.Printf("\nGet site map for %s...\n\n", url)
		for {
			// print out
			fmt.Printf("\r Found links[\033[36m%d\033[m] ", len(*c.ToArray()))
			time.Sleep(100 * time.Millisecond)
			// check if done
			if isDone {
				break
			}
		}
	}()

	// Get links now!
	e := c.Crawl(url)

	// Done! Tell spinner to stop
	isDone = true

	if e != nil {
		panic(e)
	}

	// print output

	fmt.Printf("\n\n")
	fmt.Printf("Get site map success!\n")
	fmt.Println(c.ToXml())

}
