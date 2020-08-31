package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"yourtechy.com/go-sweat/link_parser/links"
)

func TestGetHelloLinks(t *testing.T) {

	r := getReader("./data/hello.html")
	require.True(t, nil != r, "getReader failed")

	_links, err := links.GetLinks(&r)

	require.True(t, err == nil, "Get hello links returned error")
	require.True(t, _links != nil, "Get hello links failed")
	require.True(t, 1 == len(*_links), "Returned links count must be 1 item")
}

func TestGetAdjacentLinks(t *testing.T) {

	r := getReader("./data/adjacent-links.html")
	require.True(t, nil != r, "getReader failed")

	_links, err := links.GetLinks(&r)

	require.True(t, err == nil, "Get adjacent links returned error")
	require.True(t, _links != nil, "Get adjacent links failed")
	require.True(t, 2 == len(*_links), "Returned links count must be 2 items")
}

func TestGetSectionedLinks(t *testing.T) {

	r := getReader("./data/sectioned-links.html")
	require.True(t, nil != r, "getReader failed")

	_links, err := links.GetLinks(&r)

	require.True(t, err == nil, "Get sectioned links returned error")
	require.True(t, _links != nil, "Get sectioned links failed")
	require.True(t, 3 == len(*_links), "Returned links count must be 3 items")
}

func TestGetNestedLinks(t *testing.T) {

	r := getReader("./data/nested-links.html")
	require.True(t, nil != r, "getReader failed")

	_links, err := links.GetLinks(&r)

	require.True(t, err == nil, "Get nested links returned error")
	require.True(t, _links != nil, "Get nested links failed")
	require.True(t, 2 == len(*_links), "Returned links count must be 2 items")
}
