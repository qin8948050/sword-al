package main

import (
	"fmt"
)

func main() {
	MaxRange()
}

func MaxRange() {
	start:=0
	end:=3
	numbers:= []int{1,2,3,4}
	result,p,q:=maxRange(numbers,start,end)
	fmt.Println(result,p,q)
}

func maxRange(numbers []int,start,end int)(int,int,int) {
	if start<0 ||end<0 ||start>end {
		return -1,-1,0
	}
	if start==end{
		return start,end,numbers[start]
	}
	mid:=(start+end)/2
	start1,end1,sum1:=maxRange(numbers,start,mid)
	start2,end2,sum2:=maxRange(numbers,mid+1,end)
	start3,end3,total:=maxCrossingSum(numbers,start,mid,end)
	if sum1>total && sum1>sum2 {
		return start1,end1,sum1
	} else if sum2>total&& sum2>sum1 {
		return start2,end2,sum2
	} else {
		return start3,end3,total
	}
}

func maxCrossingSum(numbers []int, start, mid, end int) (int, int, int) {
	// 左半部分最大子数组和
	leftSum := numbers[mid]
	sum := 0
	leftMaxIndex := mid
	for i := mid; i >= start; i-- {
		sum += numbers[i]
		if sum > leftSum {
			leftSum = sum
			leftMaxIndex = i
		}
	}

	// 右半部分最大子数组和
	rightSum := numbers[mid+1]
	sum = 0
	rightMaxIndex := mid + 1
	for i := mid + 1; i <= end; i++ {
		sum += numbers[i]
		if sum > rightSum {
			rightSum = sum
			rightMaxIndex = i
		}
	}

	// 返回跨越中点的最大子数组和
	return leftMaxIndex, rightMaxIndex, leftSum + rightSum
}