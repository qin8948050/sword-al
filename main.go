package main

import (
	"fmt"
	"sword-al/integrality"
)

func main() {
	data:=[]int{1,2,3,4,5,6}
	integrality.ReOrder(data)
	fmt.Println(data)
}
