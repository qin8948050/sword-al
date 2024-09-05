package main

import (
	"fmt"
)

func main() {
	result:= MoreThanHalfNum([]int{1,2,2,2,3})
	fmt.Println(result)
}

func MoreThanHalfNum(numbers []int) int{
	if len(numbers)<=0 {
		return 0
	}
	result:=numbers[0]
	times:=1
	for i:=1;i<len(numbers);i++ {
		if numbers[i]==result{
			times++
		} else if times>0 {
			times--
		} else {
			result=numbers[i]
		}
	}
	return result
}