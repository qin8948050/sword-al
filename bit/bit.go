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

func IsExponseOfTwo(n int) bool {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count == 1
}