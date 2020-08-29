package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"yourtechy.com/go-sweat/link_parser/links"
)

func TestGetHelloLinks(t *testing.T) {

	_links, err := links.GetLinks("./data/hello.html")

	assert.True(t, err == nil, "Get hello links returned error")
	assert.True(t, _links != nil, "Get hello links failed")
	assert.True(t, 1 == len(*_links), "Returned links count must be 1 item")
}

func TestGetAdjacentLinks(t *testing.T) {

	_links, err := links.GetLinks("./data/adjacent-links.html")

	assert.True(t, err == nil, "Get adjacent links returned error")
	assert.True(t, _links != nil, "Get adjacent links failed")
	assert.True(t, 2 == len(*_links), "Returned links count must be 2 items")
}

func TestGetSectionedLinks(t *testing.T) {

	_links, err := links.GetLinks("./data/sectioned-links.html")

	assert.True(t, err == nil, "Get sectioned links returned error")
	assert.True(t, _links != nil, "Get sectioned links failed")
	assert.True(t, 3 == len(*_links), "Returned links count must be 3 items")
}

func TestGetNestedLinks(t *testing.T) {

	_links, err := links.GetLinks("./data/nested-links.html")

	assert.True(t, err == nil, "Get nested links returned error")
	assert.True(t, _links != nil, "Get nested links failed")
	assert.True(t, 2 == len(*_links), "Returned links count must be 2 items")
}
