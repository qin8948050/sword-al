package main

import (
	"fmt"
	"math"
)

func main() {
	result := MaxStr("babad")
	fmt.Println(result)
}

func MaxStr(s string) string {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := match(s, i, i)
		len2 := match(s, i, i+1)
		maxLength := int(math.Max(float64(len1), float64(len2)))
		if maxLength > end-start {
			// maxLength-1表示去掉中心位置后的左右长度
			start = i - (maxLength-1)/2
			end = start + maxLength - 1
		}
	}
	return s[start : end+1]
}

func match(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start - 1
}
