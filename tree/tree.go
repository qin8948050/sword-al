package tree

import "fmt"

func Image(root *TreeNode) {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		Image(root.Left)
		Image(root.Right)
	}
}

func ValidatePostOrder(sequeues []int, start, end int) bool {
	if start > end || start < 0 || end > len(sequeues) {
		return false
	}
	if start == end {
		return true
	}
	i := start
	root := sequeues[end]
	for i < end && sequeues[i] < root {
		i++
	}

	for j := i; j < end; j++ {
		if sequeues[j] < root {
			return false
		}
	}
	return ValidatePostOrder(sequeues, start, i-1) && ValidatePostOrder(sequeues, i, end-1)
}

func FindPath(root *TreeNode, expected int) {
	stacks:=make([]*TreeNode,0)
	current:=0
	findPath(root,stacks,&current,expected)
}

func findPath(node *TreeNode, stacks []*TreeNode, current *int,expected int) {
	if node == nil {
		return
	}
	*current += node.Value
	stacks = append(stacks, node)
	if node.Left == nil && node.Right == nil && *current==expected{
		fmt.Println("a path found,path:")
		for _,node =range stacks{
			fmt.Println(node.Value)
		}
	}
	if node.Left!=nil{
		findPath(node.Left,stacks,current,expected)
	}
	if node.Right!=nil{
		findPath(node.Right,stacks,current,expected)
	}
	*current-=node.Value
	stacks=stacks[:len(stacks)-1]
}