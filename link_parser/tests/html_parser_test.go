package tests

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	parser "yourtechy.com/go-sweat/link_parser/parser/html"
)

const (
	helloTest   = "./data/hello.html"
	missingFile = "./data/missing.html"
)

func getReader(source string) io.Reader {

	r, err := os.Open(source)

	if err != nil {
		fmt.Println("os.Open failed!", err)
		return nil
	}

	return io.Reader(r)
}

func TestCreateHtmlParserMissinFile(t *testing.T) {

	r := getReader(missingFile)
	require.True(t, nil == r, "file should not exist")

}

func TestLoadHelloTestFile(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")
	require.True(t, nil != p, "NewHtmlParser failed to create parser")
}

func TestGetRootParentReturnsNil(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")

	parent := p.GetParent()

	require.True(t, nil == parent, "GetParent should return nil")
}

func TestGetNode(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")

	node := p.FindNode("body")

	require.True(t, nil != node, "GetNode should return body element")
}

func TestGetParent(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")

	child := p.GetChild() // body node

	require.True(t, nil != child, "GetChild body returned nil")

	parent := child.GetParent() // html node

	require.True(t, nil != parent, "GetParent returned nil")
}

func TestGetSibling(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")

	child := p.GetChild() // body node

	require.True(t, nil != child, "GetChild body returned nil")

	child = child.GetChild() // get body first child element

	require.True(t, nil != child, "Body GetChild returned nil")

	var count int = 1

	for {
		child = child.NextSibling()
		if child != nil {
			count++
		} else {
			break
		}
	}

	require.True(t, count == 2, "Body should have 2 children")
}

func TestGetChildNode(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")
	require.True(t, nil != p, "NewHtmlParser returned nil")

	child := p.GetChild()

	require.True(t, nil != child, "GetChild returned nil")
}

func TestGetNodeTextRecursive(t *testing.T) {

	r := getReader(helloTest)
	require.True(t, nil != r, "getReader failed")

	p, err := parser.NewHtmlParser(&r)

	require.True(t, nil == err, "NewHtmlParser returned error")
	require.True(t, nil != p, "NewHtmlParser returned nil")

	text := p.GetText(0)

	require.True(t, "" != text, "GetText returned empty string")
}
