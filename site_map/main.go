package main

import (
	"fmt"
	"time"

	"github.com/tj/go-spin"
	"yourtechy.com/go-sweat/site_map/crawler"
)

func main() {

	var isDone bool = false
	var url = "https://yourtechy.com"

	// create progress spinner

	s := spin.New()
	s.Set(spin.Box1)

	// run spinner in background
	fmt.Println()
	go func() {
		for {
			// print progress wheel
			fmt.Printf("\rGet links for %s. \033[36mProgress:\033[m %s ",
				url, s.Next())
			time.Sleep(100 * time.Millisecond)
			// check if done
			if isDone {
				fmt.Printf("\n\n")
				break
			}
		}
	}()

	// Get links now!

	c := crawler.NewCrawler()
	e := c.Crawl(url)

	// Done! Tell spinner to stop
	isDone = true

	if e != nil {
		panic(e)
	}

	// print output

	fmt.Printf("Get links success!\n")
	fmt.Println(c.ToXml())

}
