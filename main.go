package main

import (
	"fmt"
)

func main() {
	result,ok := MoreThanHalfNum([]int{1,2,2,2,3},5)
	fmt.Println(result,ok)
}

