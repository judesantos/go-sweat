package parser

type INode interface {
	Parent() *INode
	NextSibling() *INode
	NextChild() *INode
}
