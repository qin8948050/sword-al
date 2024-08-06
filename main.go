package main

import (
	"sword-al/tree"
)

func main() {
	root:=tree.InitTree()
	tree.PostTraverse(root)
}
