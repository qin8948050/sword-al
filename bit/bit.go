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