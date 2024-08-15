package tree

func Image(root *TreeNode) {
	if root!=nil{
		root.Left,root.Right=root.Right,root.Left
		Image(root.Left)
		Image(root.Right)
	}
}

func ValidatePostOrder(sequeues []int,start,end int) bool {
	if start>end || start<0 || end>len(sequeues){
		return false
	}
	if start==end {
		return true
	}
	i:=start
	root:=sequeues[end]
	for i<end && sequeues[i]<root{
		i++
	}

	for j:=i;j<end;j++{
		if sequeues[j]<root {
			return false
		}
	}
	return ValidatePostOrder(sequeues,start,i-1) && ValidatePostOrder(sequeues,i,end-1)
}