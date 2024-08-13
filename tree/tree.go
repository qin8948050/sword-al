package tree

func Image(root *TreeNode) {
	if root!=nil{
		root.Left,root.Right=root.Right,root.Left
		Image(root.Left)
		Image(root.Right)
	}
}