package main

import (
	"fmt"
	"sword-al/tree"
)

func main() {
	// p6:=&integrality.ListNode{
	// 	Value: 6,
	// 	Next: nil,
	// }
	// p5:=&integrality.ListNode{
	// 	Value: 5,
	// 	Next: p6,
	// }
	// p4:=&integrality.ListNode{
	// 	Value: 4,
	// 	Next: p5,
	// }
	// p3:=&integrality.ListNode{
	// 	Value: 3,
	// 	Next: p4,
	// }
	// p2:=&integrality.ListNode{
	// 	Value: 2,
	// 	Next: nil,
	// }
	// p1:=&integrality.ListNode{
	// 	Value: 1,
	// 	Next: p2,
	// }
	// p22:=&integrality.ListNode{
	// 	Value: 2,
	// 	Next: nil,
	// }
	// p11:=&integrality.ListNode{
	// 	Value: 1,
	// 	Next: p22,
	// }
	// result:=integrality.Merge(p1,p11)
	// fmt.Println(result)
	a:=tree.TreeNode{Left:&tree.TreeNode{Value: 2},Right:nil,Value:1}
	// b:=tree.TreeNode{Left:nil,Right:nil,Value:1}
	result:=tree.HasSubTree(&a,nil)
	fmt.Println(result)
}
