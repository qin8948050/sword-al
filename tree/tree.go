package tree

import (
	"fmt"
)

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

func TreeToListNode(root *TreeNode) *TreeNode{
	if root==nil ||(root.Left==nil&& root.Right==nil){
		return root
	}
	slice:=make([]*TreeNode,0)
	TraverseTreeNode(root,&slice)
	pre:=0
	post:=pre+1
	for post<len(slice) {
	  slice[pre].Right=slice[post]
	  pre++
	  post++
	}
	post=len(slice)-1
	pre=post-1
	for pre>=0 {
	  slice[post].Left=slice[pre]
	  post--
	  pre--
	}
	return slice[0]
}

func TraverseTreeNode(root *TreeNode,slices *[]*TreeNode){
	if root!=nil{
		TraverseTreeNode(root.Left,slices)
		*slices=append(*slices, root)
		TraverseTreeNode(root.Right,slices)
	}

}

func Spin(root *TreeNode) {
	if root==nil {
		return 
	}
	currentStack:=[]*TreeNode{root}
	nextStack:=make([]*TreeNode,0)
	level:=1
	for len(currentStack)>0 {
		node:=currentStack[len(currentStack)-1]
		currentStack=currentStack[:len(currentStack)-1]
		fmt.Println(node.Value)
		if level%2==0 {
			if node.Left!=nil{
				nextStack=append(nextStack, node.Left)
			}
			if node.Right!=nil{
				nextStack=append(nextStack, node.Right)
			}
		} else {
			if node.Right!=nil{
				nextStack=append(nextStack, node.Right)
			}
			if node.Left!=nil{
				nextStack=append(nextStack, node.Left)
			}
		}
		if len(currentStack)==0 {
			level++
			currentStack,nextStack=nextStack,currentStack
			nextStack=nil
		}
	}
}

func MaxDepth(root *TreeNode) int{
	if root==nil{
		return 0
	}
	leftDepth:=MaxDepth(root.Left)
	rightDepth:=MaxDepth(root.Right)
	if leftDepth>rightDepth {
		return leftDepth+1
	}  else {
		return rightDepth+1
	}
}

// 搜索二叉树的第k大元素
func KMaxNode(root *TreeNode,k int) (int,bool) {
	var result *TreeNode
	kMaxNode(root,&result,&k)
	if result==nil{
		return 0,false
	}
	return result.Value,true
}


 func kMaxNode(root *TreeNode,result **TreeNode,k *int){
	if root==nil {
		return
	}
	kMaxNode(root.Right,result,k)
	*k--;
	if *k==0 {
		*result=root
	}
	kMaxNode(root.Left,result,k)
 }