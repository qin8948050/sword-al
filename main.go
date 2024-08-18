package main

import (
	"fmt"
	"sword-al/tree"
)

func main() {
	node:=&tree.TreeNode{
		Left: &tree.TreeNode{
			Value:9,
			Left: &tree.TreeNode{
				Value: 7,
			},
			Right: &tree.TreeNode{
				Value: 10,
			},
		},
		Value: 15,
		Right: &tree.TreeNode{
			Left: &tree.TreeNode{
				Value: 16,
			},
			Right: &tree.TreeNode{
				Value: 19,
			},
			Value: 18,
		},
	}
	result,ok:=tree.KMaxNode(node,2)
	fmt.Println(result,ok)
}



