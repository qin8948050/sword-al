package main

import (
	"fmt"
)

func main() {
	a:=[]int{5,1,2,3,4}
	partitionByRatio(a,0,len(a)-1,2)
	fmt.Println(a[2])
}

func partitionByRatio(numbers []int,low int,high int,target int){
	if low<high{
		leftSubEnd:=partition(numbers,low,high)
		if leftSubEnd==target{
			return
		}
		if leftSubEnd<target {
			partitionByRatio(numbers,leftSubEnd+1,high,target)
		} else {
			partitionByRatio(numbers,low,leftSubEnd-1,target)
		}
	}
}

func partition(numbers []int,low,high int) int {
	pivlot:=numbers[high]
	i:=low
	for j:=low;j<high;j++{
		if numbers[j]<pivlot{
			numbers[j],numbers[i]=numbers[i],numbers[j]
			i++
		}
	}
	numbers[i],numbers[high]=numbers[high],numbers[i]	
	return i
}

