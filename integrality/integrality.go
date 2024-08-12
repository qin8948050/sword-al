package integrality

import (
	"math"
)

func Equal(a,b float64) bool {
	return math.Abs(a-b)<1e-9 
}

var InvalidInput bool 
func Power(base float64,exponse int) float64 {
	//排除底数为0，指数为负数的情况
	if Equal(base,0.0) && exponse<0 {
		InvalidInput=true
		return 0
	}
	absExponse:=exponse
	if exponse<0{
		absExponse=-exponse
	}
	result:=powerExponse(base,int(absExponse))
	if exponse<0{
		return 1/result
	}
	return result
}

func powerExponse(base float64,exponse int) float64 {
	var result float64=1
	for i:=1;i<=exponse;i++{
		result*=base
	}
	return result
}

type ListNode struct {
	Value int
	Next *ListNode
}
// O(1)删除链表的节点
func DeleteNode(head *ListNode,toBeDeleted *ListNode) {
	if head==nil || toBeDeleted==nil{
		return
	}
	if toBeDeleted.Next!=nil{
		pNext:=toBeDeleted.Next
		toBeDeleted.Value=pNext.Value
		toBeDeleted.Next=pNext.Next
		pNext=nil

	} else if head==toBeDeleted{
		toBeDeleted=nil
	} else {
		pNode:=head
		for pNode!=nil{
			if pNode.Next==toBeDeleted{
				pNode.Next=toBeDeleted.Next
				toBeDeleted=nil
			}
			pNode=pNode.Next
		}
	}
}


// 将数组上的奇数移动到偶数前
func ReOrder(data []int,judge func(d int) bool) {
	if len(data)==0 {
		return
	}
	head:=0
	tail:=len(data)-1
	for head<tail {

		for head<tail && !judge(data[head]){
			head++
		}
		for head<tail && judge(data[tail]){
			tail--
		}
		if head<tail {
			temp:=data[tail]
			data[tail]=data[head]
			data[head]=temp
		}
	}
}

// 倒数第K个节点
func FindTheKNode(head *ListNode,k int) *ListNode {
	if head==nil ||k<=0{
		return nil
	}
	p1:=head

	for i:=0;i<k;i++{
		if p1==nil{
			return nil
		}
		p1=p1.Next
	}
	p2:=head
	for p1!=nil{
		p1=p1.Next
		p2=p2.Next
	}
	return p2
}

func ReverseList(head *ListNode) *ListNode{
	current:=head
	var pReversedHead *ListNode
	var frontend *ListNode
	for current!=nil{
		tail:=current.Next
		if tail==nil{
			pReversedHead=current
		}
		current.Next=frontend
		frontend=current
		current=tail
	}
	return pReversedHead
}

func Merge(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1==nil {
		return l2
	}
	if l2==nil{
		return l1
	}
	head:=new(ListNode)
	if l1.Value<l2.Value{
		head=l1
		head.Next=Merge(l1.Next,l2)
	}
}