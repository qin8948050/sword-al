package main

import (
	"sword-al/analysis"
)

func main() {
	matrix:=[][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}
	analysis.PrintMatrixInCircle(matrix,4,4)
}
