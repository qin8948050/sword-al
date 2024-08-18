package main

import (
	"sword-al/tree"
)

func main() {
    root := &tree.TreeNode{Value: 1}
    root.Left = &tree.TreeNode{Value: 2}
    root.Right = &tree.TreeNode{Value: 3}
    root.Left.Left = &tree.TreeNode{Value: 4}
    root.Left.Right = &tree.TreeNode{Value: 5}
    root.Right.Left = &tree.TreeNode{Value: 6}
    root.Right.Right = &tree.TreeNode{Value: 7}
	tree.Spin(root)
}



