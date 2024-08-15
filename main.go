package main

import (
	"sword-al/tree"
)

func main() {
	node:=&tree.TreeNode{
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
	}
	tree.FindPath(node,22)
}



