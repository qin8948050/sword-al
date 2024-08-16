package main

import (
	"fmt"
	"sword-al/tree"
)

func main() {
/* 	node:=&tree.TreeNode{
		Left: &tree.TreeNode{
			Left: &tree.TreeNode{
				Left: nil,
				Right: nil,
				Value: 4,
			},
			Right:&tree.TreeNode{
				Left: nil,
				Right: nil,
				Value: 7,
			},
			Value:5,
		},
		Right: &tree.TreeNode{
			Left: nil,
			Right: nil,
			Value: 12,
		},
		Value: 10,
	} */
	node:=&tree.TreeNode{
		Left: nil,
		Right: nil,
		Value: 1,
	}
	fmt.Println(tree.TreeToListNode(node))
}



