package bit

import "strings"

func ExcelColumnToNumber(columnLabel string) int {
	columnLabel = strings.ToUpper(columnLabel) // 处理大小写不敏感
	columnNumber := 0

	for _, char := range columnLabel {
		// 计算每个字符对应的数字值（A=1, B=2, ..., Z=26）
		value := int(char - 'A' + 1)
		// 累加到总的列数
		columnNumber = columnNumber*26 + value
	}

	return columnNumber
}

// 计算一个整数的二进制表示中，1的个数
// 一个整数减去1后的结果，再与原整数做与运算，相等于将整数二进制的最右边1变为0
func CountOnesInBinaryCount(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

// 用来检查一个整数 n 是否是二的幂
// 2的幂，二进制表示只有1个1
// n-1: n的唯一的1位变为 0，其余的0位变为1
// n&(n-1) 结果为0
func IsExponseOfTwo(n int) bool {
	return n>0&& (n&(n-1))==0
}

// 数组中只出现一次的两个数，其他都出现两次
func FindTwoSingleNumbers(numbers []int) (int,int) {
   xor:=0
   //1.计算整体异或结果，重复的数字会相互抵消，最后只剩出现一次数的异或结果
   for _,n:=range numbers {
      xor^=n
   }
   //2.找差异位，不同组的数分为出现在1和0的不同位置
   diff:=xor&(-xor)
   num1:=0
   num2:=0
   //3.分组
   for _,n:=range numbers {
      //1组
      if n & diff==1 {
         num1^=n
      } else {
         num2^=n
      }
   }
   return num1,num2
}