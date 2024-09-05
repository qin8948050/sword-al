package sort


// 数字旋转，求最小值。
func MinNum(input []int) int{
	if len(input)==0 {
		panic("input isn't empty")
	}
	head:=0
	tail:=len(input)-1
	mid:=head
	for(input[head]>=input[tail]){
		if tail-head==1 {
			mid=tail
			break
		}
		mid=(head+tail)/2
		if input[head]==input[tail] && input[head]==input[mid] {
			return minInOrder(input,head,tail)
		}
		if input[mid]>=input[head] {
			head=mid
		} else if input[mid]<=input[tail] {
			tail=mid
		}
	}
	return input[mid]
}

func minInOrder(input []int,head int,tail int) int{
	var result int
	for head<=tail {
		if input[head]<=result{
			result=input[head]
		}
		head++
	}
	return result
}

func Add(n int) int {
	if n==0 {
		return 0
	}
	if n==1 {
		return 1
	}
	return n+Add(n-1)
}

func Fib(n int) int{
	if n<=1 {
		return n
	}
	a:=0
	b:=1
	var c int
	for i:=2;i<=n;i++{
		c=b
		b=a+b
		a=c
	}
	return b
}

func FibRecur(n int) int{
	if n<=1 {
		return n
	}
	return FibRecur(n-1)+FibRecur(n-2)
}

func BubbleSort(data []int) {
	for i:=0;i<len(data);i++ {
	for j:=1;j<len(data)-i;j++{
		if data[j-1]>data[j] {
			data[j-1],data[j]=data[j],data[j-1]
		}
	}
	}
}

// 快速排序
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
	//放置枢纽
	numbers[i],numbers[high]=pivlot,numbers[i]	
	return i
}

//查找中位数
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

// 个数超过一般的元素
//方法1：分治
func MoreThanHalfNum(numbers []int,length int) (int,bool){
	if !checkNumbersValid(numbers) {
		return 0,false
	}
	mid:=length>>1
	start:=0
	end:=length-1
	index:=partition(numbers,start,end)
	for index!=mid {
		if index>mid {
			end=index-1
		} else {
			start=index+1
		}
		index=partition(numbers,start,end)
	}
	midNumber:=numbers[mid]
	if checkThanHalfNum(numbers,length,midNumber) {
		return midNumber,true
	}
	return 0,false

}

// 方法2：出现次数超过一般的元素，说明比其他元素的次数加起来都多
func MoreThanHalfNum1(numbers []int) (int,bool){
	if !checkNumbersValid(numbers) {
		return 0,false
	}
	length:=len(numbers)
	result:=numbers[0]
	times:=1
	for i:=1;i<length;i++ {
		if numbers[i]==result{
			times++
		} else if times>0 {
			times--
		} else {
			result=numbers[i]
		}
	}
	if checkThanHalfNum(numbers,length,result) {
		return result,true
	}
	return 0,false
}
func checkNumbersValid(numbers []int) bool {
	if numbers==nil ||len(numbers)<=0 {
		return false
	}
	return true
}



func checkThanHalfNum(numbers []int,length int,midNumber int) bool {
	var times int
	for i:=0;i<length;i++ {
		if numbers[i]==midNumber {
			times++
		}
	}
	return times>length/2
}