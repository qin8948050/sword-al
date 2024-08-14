package main

import (
	"fmt"
	"sword-al/stack"
)

func main() {
	result:=stack.VlidateStackSequences([]int{1,2,3,4,5},[]int{5,4,3,1,2})
	fmt.Println(result)
}



