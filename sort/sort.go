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



// 最大子序列和，方法1
func MaxSubArraySum(numbers []int) ([]int,int) {
	if len(numbers)<=0{
		return nil,0
	}
	maxSum:=0
	currentSum:=0
	start:=0
	end:=0
	for 	i:=0;i<len(numbers);i++{
		if currentSum<=0 {
			currentSum=numbers[i]
			start=i
		} else {
			currentSum+=numbers[i]
		}
		if currentSum>maxSum{
			end=i
			maxSum=currentSum
		}
	}
	return numbers[start:end+1],maxSum
}

//  最大子序列和,方法2
func maxSubArraySum1(numbers []int) int {
	dp:=make([]int,len(numbers))
	dp[0]=numbers[0]
	maxResult:=0
	for i:=1;i<len(numbers);i++ {
		if dp[i-1]<0 {
			dp[i]=numbers[i]
		} else {
			dp[i]=dp[i-1]+numbers[i]
		}
		if maxResult<dp[i] {
			maxResult=dp[i]
		}
	}
	return maxResult
}



// 获取丑数的函数,动态规划，空间换时间
func GetUglyNumber_Solution2(index int) int {
    if index <= 0 {
        return 0
    }

    // 创建一个数组来存储丑数
    uglyNumbers := make([]int, index)
    uglyNumbers[0] = 1
    nextUglyIndex := 1

    // 三个指针分别指向当前需要乘以2，3，5的最小丑数
    pMultiply2 := 0
    pMultiply3 := 0
    pMultiply5 := 0

    // 循环生成丑数
    for nextUglyIndex < index {
        min:=Min(uglyNumbers[pMultiply2]*2,uglyNumbers[pMultiply3]*3,uglyNumbers[pMultiply5]*5)
        uglyNumbers[nextUglyIndex]=min
        
        if uglyNumbers[pMultiply2]*2==min{
            pMultiply2++
        }
        if uglyNumbers[pMultiply3]*3==min{
            pMultiply3++
        }
        if uglyNumbers[pMultiply5]*5==min{
            pMultiply5++
        }
        nextUglyIndex++
    }

    // 返回生成的丑数
    return uglyNumbers[nextUglyIndex-1]
}

// 找最小值的辅助函数
func Min(a, b, c int) int {
    min := a
    if b < min {
        min = b
    }
    if c < min {
        min = c
    }
    return min
}


//逆序组数量：归并排序，分治思想
func CountOfInversePair(numbers []int,start,end int) int {
   if start>=end{
      return 0
   }
   mid:=(start+end)/2
   //左、右序列逆序组数量
   count:=CountOfInversePair(numbers,start,mid)+CountOfInversePair(numbers,mid+1,end)
   //跨左右序列逆序组数量
   count+=Merge(numbers,start,mid,end)
   return count
}

func Merge(numbers []int,start,mid,end int)int {
   if start>=end {
      return 0
   }
   i:=start
   j:=mid+1
   arrs:=make([]int,0)
   count:=0
   for i<=mid&&j<=end {
      if numbers[i]<numbers[j] {
         arrs = append(arrs, numbers[i])
         i++
      } else {
         arrs = append(arrs, numbers[j])
	   //计算逆序组数量，左子数组是排序好的，因此从 i 到 mid 的所有元素都比 arr[j] 大，这些元素都与 arr[j] 形成逆序对
         count+=(mid-i+1)
         j++
      }
   }
   arrs=append(arrs, numbers[i:mid+1]...)
   arrs=append(arrs, numbers[j:end+1]...)
   k:=0
   for i:=start;i<=end;i++ {
      numbers[i]=arrs[k]
      k++
   }
   return count
}