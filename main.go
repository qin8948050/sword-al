package main

import (
	"fmt"
	"sword-al/tree"
)

func main() {
	result:=tree.ValidatePostOrder([]int{1,2},0,0)
	fmt.Println(result)
}



