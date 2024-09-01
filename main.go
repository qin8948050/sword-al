package main

import (
	"fmt"
	"math"
)

func main() {
	a:=[][]int{
		{2,10},
		{0,5},
		{8,10},
	}
	result:=Merge(a)
	fmt.Println(result)
}

func BubbleSort(data [][]int) {
	for i:=0;i<len(data)-1;i++ {
		for j:=0;j<len(data)-1-i;j++ {
			if  data[j][0]>data[j+1][0] || (data[j][0]==data[j][0] && data[j][1]>data[j+1][1]){
				data[j],data[j+1]=data[j+1],data[j]
			}
		}
	}
}

func Merge(ranges [][]int) [][]int{
	if len(ranges)<=1 {
		return ranges
	}
	BubbleSort(ranges)
	result:=[][]int{{ranges[0][0],ranges[0][1]}}
	for i:=1;i<len(ranges);{
		l1:=result[len(result)-1][0]
		u1:=result[len(result)-1][1]
		l2:=ranges[i][0]
		u2:=ranges[i][1]
		if u1>=l2{
			sub:=[]int{l1,int(math.Max(float64(u1),float64(u2)))}
			if len(result)>0 {
				result=result[:len(result)-1]
			}
			result=append(result, sub)
		}else {
			result=append(result, []int{l1,u1},[]int{l2,u2})
		}
		i++
	}
	return result
}

