package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	parser "yourtechy.com/go-sweat/link_parser/parser/html"
)

const (
	helloTest   = "./data/hello.html"
	missingFile = "./data/missing.html"
)

func TestCreateHtmlParserMissingFile(t *testing.T) {

	_, err := parser.NewHtmlParser(missingFile)

	assert.True(t, nil != err, "NewHtmlParser should return error")
}

func TestLoadHelloTestFile(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")
	assert.True(t, nil != p, "NewHtmlParser failed to create parser")
}

func TestGetRootParentReturnsNil(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")

	parent := p.GetParent()

	assert.True(t, nil == parent, "GetParent should return nil")
}

func TestGetNode(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")

	node := p.FindNode("body")

	assert.True(t, nil != node, "GetNode should return body element")
}

func TestGetParent(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")

	child := p.GetChild() // body node

	assert.True(t, nil != child, "GetChild body returned nil")

	parent := child.GetParent() // html node

	assert.True(t, nil != parent, "GetParent returned nil")
}

func TestGetSibling(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")

	child := p.GetChild() // body node

	assert.True(t, nil != child, "GetChild body returned nil")

	child = child.GetChild() // get body first child element

	assert.True(t, nil != child, "Body GetChild returned nil")

	var count int = 1

	for {
		child = child.NextSibling()
		if child != nil {
			count++
		} else {
			break
		}
	}

	assert.True(t, count == 2, "Body should have 2 children")
}

func TestGetChildNode(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")
	assert.True(t, nil != p, "NewHtmlParser returned nil")

	child := p.GetChild()

	assert.True(t, nil != child, "GetChild returned nil")
}

func TestGetNodeTextRecursive(t *testing.T) {

	p, err := parser.NewHtmlParser(helloTest)

	assert.True(t, nil == err, "NewHtmlParser returned error")
	assert.True(t, nil != p, "NewHtmlParser returned nil")

	text := p.GetText(0)

	assert.True(t, "" != text, "GetText returned empty string")
}
