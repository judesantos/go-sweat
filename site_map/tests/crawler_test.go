package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"yourtechy.com/go-sweat/site_map/crawler"
)

const (
	siteUrl = "https://yourtechy.com"
)

func TestCreateCrawler(t *testing.T) {

	c := crawler.NewCrawler()
	err := c.Crawl(siteUrl)

	assert.True(t, nil == err, "GetLinks returned with error")
	assert.True(t, 0 < len(*c.ToArray()), "GetLinks returned empty links")
}
