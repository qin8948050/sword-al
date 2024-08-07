package main

import (
	"fmt"
	"sword-al/stack"
)

func main() {
	q:=stack.Stack{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
