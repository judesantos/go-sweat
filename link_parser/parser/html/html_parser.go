package parser

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"yourtechy.com/go-sweat/utils/logger"
)

var (
	log = logger.NewLogger()
)

/************************************************************************
 *
 * Public methods and types
 *
 ***********************************************************************/

type HtmlAttribute struct {
	Name  string
	Value string
}

// HtmlNode - App wrapper to node provider (x/net/html)
type HtmlNode struct {
	iNode *html.Node
}

type INode interface {
	Parent() *INode
	NextSibling() *INode
	NextChild() *INode
}

// GetAtribute - Get this node attribute
func (node *HtmlNode) GetAttribute(key string) *HtmlAttribute {
	if node != nil {
		for _, attr := range node.iNode.Attr {
			if attr.Key == key {
				return &HtmlAttribute{
					Name:  key,
					Value: attr.Val,
				}
			}
		}
	}

	return nil
}

func (node *HtmlNode) GetNodeName() string {

	if node != nil {
		return string(node.iNode.Data)
	}

	return ""
}

// NewHtmlParser - Create new Html parser
// @return the HTMLNode instance
func NewHtmlParser(source string) (*HtmlNode, error) {

	f, err := os.Open(source)
	if err != nil {
		log.Error(fmt.Sprintf("open %s failed!", source), err)
		return nil, err
	}

	node, err := html.Parse(f)
	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}

	return &HtmlNode{
		iNode: node,
	}, nil
}

func (node *HtmlNode) GetType() uint32 {
	return uint32(node.iNode.Type)
}

// Parent - Return parent node
func (node *HtmlNode) GetParent() *HtmlNode {

	if node.iNode == nil {
		return nil
	}

	if node.iNode.Parent == nil {
		return nil
	}

	return &HtmlNode{
		iNode: node.iNode.Parent,
	}
}

// GetChild - Get instance of child of this node
func (node *HtmlNode) GetChild() *HtmlNode {

	if node.iNode == nil {
		return nil
	}

	if node.iNode.FirstChild == nil {
		return nil
	}

	return &HtmlNode{
		iNode: node.iNode.FirstChild,
	}
}

// NextSibling - Get instance of the next sibling of this node
func (node *HtmlNode) NextSibling() *HtmlNode {

	if node.iNode == nil {
		return nil
	}

	if node.iNode.NextSibling == nil {
		return nil
	}

	if node.iNode != nil {
		node.iNode = node.iNode.NextSibling
	}

	return &HtmlNode{
		iNode: node.iNode,
	}
}

// FindNode - Return first found occurrence of the html element
//           as specified by 'elemName'
func (node *HtmlNode) FindNode(elemName string) *HtmlNode {

	n := searchNode(node.iNode, elemName)

	if n != nil {
		return &HtmlNode{
			iNode: n,
		}
	}

	return nil
}

/*
GetText - Get text data.
Recursively extract strings from children,
and descendants as prescribed by
param 'depth'.
***
If 'depth' is zero, iterate through this
node's descendants.

@return this node's text data
*/
func (node *HtmlNode) GetText(depth int) string {
	return recursiveGetText(node.iNode, 0, depth)
}

/************************************************************************
 *
 * Internal methods
 *
 ***********************************************************************/

// recursiveGetText - Iterate node and extract all element string data, if any.
func recursiveGetText(
	node *html.Node,
	currDepth int,
	targetDepth int,
) string {

	var text string

	if targetDepth > 0 && currDepth == targetDepth {
		return ""
	}

	if node.Type == html.TextNode {
		text = node.Data
	} else {
		for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
			text += recursiveGetText(ch, currDepth+1, targetDepth)
		}
	}

	return strings.TrimSpace(text)
}

// searchNode - Iterate node, find first occurrence of element
//              specified by element name in param 'target'
func searchNode(n *html.Node, target string) *html.Node {

	var ch *html.Node

	if target != "" {

		if n.Type == html.ElementNode && n.Data == target {
			return n
		}

		// see if we get something from children

		if ch = n.FirstChild; ch != nil {
			ch = searchNode(ch, target)
		}

		if ch != nil {
			return ch
		}

		// not found in child, let's try adjacent nodes

		if ch = n.NextSibling; ch != nil {
			ch = searchNode(ch, target)
		}
	}

	return ch
}
