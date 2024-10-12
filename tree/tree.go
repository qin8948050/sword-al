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


// 判断是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
    return checkDepth(root) != -1
}

func checkDepth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    

    leftDepth := checkDepth(node.Left)
    rightDepth := checkDepth(node.Right)

    if abs(rightDepth-leftDepth)>1 {
        return -1
    }
    return max(leftDepth, rightDepth) + 1
}

// 返回两个整数的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 返回整数的绝对值
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}



// 递归查找两个节点的最低公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点是 nil，或者当前节点等于 p 或 q，则直接返回当前节点
	if root == nil || root == p || root == q {
		return root
	}

	// 递归在左子树和右子树中查找
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果左子树和右子树都能找到 p 或 q，说明当前节点是 LCA
	if left != nil && right != nil {
		return root
	}

	// 如果只有一侧找到 p 或 q，返回该侧的结果
	if left != nil {
		return left
	}
	return right
}

// 非递归查找两个节点的最低公共祖先
func findPath1(root *TreeNode,paths *[]*TreeNode,target *TreeNode) bool {
   if root==nil {
      return false
   }
   *paths = append(*paths, root)

   if root==target {
      return true
   }

   if findPath1(root.Left,paths,target)||findPath1(root.Right,paths,target) {
      return true
   }
   *paths=(*paths)[:len(*paths)-1]
   return false
}

func LowestCommonAncestor1(root *TreeNode,p *TreeNode,q *TreeNode) *TreeNode {
   pathP:=make([]*TreeNode,0)
   pathQ:=make([]*TreeNode,0)
   if !findPath1(root,&pathP,p)||!findPath1(root,&pathQ,q) {
      return nil
   }

   for i:=0;i<len(pathP)&&i<len(pathQ);i++ {
      if pathP[i]==pathQ[i] {
         return pathP[i]
      }
   }
   return nil
}




type Node struct {
	Value int
	Next  *Node
}

// 链表中环的入口节点
func FirstCircleNode(head *Node) *Node {
	// 快慢指针初始化为链表头部
	slow, fast := head, head

	// 检测是否有环
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步

		// 如果快慢指针相遇，说明有环
		if slow == fast {
			break
		}
	}

	// 如果没有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 将慢指针移回链表头部，快指针保持在相遇点(慢指针从头到环入口与快指针从相遇点到环入口距离相等)
	slow = head
	// 两个指针同时走一步，直到再次相遇
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// 相遇的节点就是环的入口
	return slow
}