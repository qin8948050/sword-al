package tree

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	Data []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.Data = append(s.Data, node)
}

func (s *Stack) Empty() bool {
	if len(s.Data) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() *TreeNode {
	if len(s.Data) == 0 {
		return nil
	}
	n := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return n
}

func InitTree() *TreeNode {
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}
	root.Right.Left = &TreeNode{Value: 6}
	root.Right.Right = &TreeNode{Value: 7}
	return root
}

func PreTraverse(root *TreeNode) {
	stack := &Stack{}
	if root == nil {
		return
	}
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop()
		fmt.Println(node.Value)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

func InTraverse(root *TreeNode) {
	stack := &Stack{}

	if root == nil {
		return
	}
	node := root
	for node != nil || !stack.Empty() {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			n := stack.Pop()
			fmt.Println(n.Value)
			node = n.Right
		}
	}
}
