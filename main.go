package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,1,2,2,3,2}
	QuickSort(a,1,len(a)-1)
	fmt.Println(a)
}

func QuickSort(numbers []int,low int,high int) {
	if low < high{
		pivlot:=partition(numbers,low,high)
		QuickSort(numbers,low,pivlot-1)
		QuickSort(numbers,pivlot+1,high)
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