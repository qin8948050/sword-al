package main

import (
	"fmt"
	"sword-al/integrality"
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
	p2:=&integrality.ListNode{
		Value: 2,
		Next: nil,
	}
	p1:=&integrality.ListNode{
		Value: 1,
		Next: p2,
	}
	result:=integrality.ReverseList(p1)
	fmt.Println(*result)
}
