package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,1,2,2,3,2}
	QuickSort(a,1,len(a)-1)
	fmt.Println(a)
}

