package main

import (
	"fmt"
)

func main() {
	a:=[]int{7,1,4,3,5,5,9,4,10,25,11,12,33,2,13,6}
	result:=MaxSubArray(a)
	fmt.Println(result)
}

func MaxSubArray(numbers []int) []int {
	if len(numbers)<=0{
		return numbers
	}
	m:=make(map[int]int,0)
	for i:=0;i<len(numbers);i++ {
		currentNumber:=numbers[i]
		if _,ok:=m[currentNumber];ok{
			continue
		} else if previousValue,ok:=m[currentNumber-1];ok{
			m[currentNumber]=previousValue+1
			delete (m,currentNumber-1)
		} else {
			m[currentNumber]=1
		}
	}
	maxCount:=0
	var maxKey int
    for k,v:=range m {
		if v>maxCount{
			maxCount=v
			maxKey=k
		}
	}

	result:=make([]int,maxCount)
	for i:=maxCount-1;i>=0;i-- {
		result[i]=maxKey
		maxKey--
	}
	return result
}

