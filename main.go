package main

import (
	"fmt"
	"sword-al/sort"
)

func main() {
	input:=[]int{1,1,1,0,1}
	result:=sort.MinNum(input)
	fmt.Println(result)
}
