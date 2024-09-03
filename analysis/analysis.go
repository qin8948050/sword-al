package analysis

import (
	"fmt"
	"math"
)

func PrintMatrixInCircle(matrix [][]int, rows, columns int) {
	if rows == 0 && columns == 0 {
		return
	}
	start := 0
	for rows > start*2 && columns > start*2 {
		printMatrixInCircle(matrix, rows, columns, start)
		start++
	}
}

func printMatrixInCircle(matrix [][]int, rows, columns, start int) {
	endX := columns - start - 1
	endY := rows - start - 1
	for i := start; i <= endX; i++ {
		fmt.Println(matrix[start][i])
	}
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			fmt.Println(matrix[i][endX])
		}
	}
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			fmt.Println(matrix[endY][i])
		}
	}
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
			fmt.Println(matrix[i][start])
		}
	}
}

// 最长子序列
func MaxSubArray(numbers []int) []int {
	if len(numbers) <= 0 {
		return numbers
	}
	m := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		currentNumber := numbers[i]
		if _, ok := m[currentNumber]; ok {
			continue
		} else if previousValue, ok := m[currentNumber-1]; ok {
			m[currentNumber] = previousValue + 1
			delete(m, currentNumber-1)
		} else {
			m[currentNumber] = 1
		}
	}
	maxCount := 0
	var maxKey int
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			maxKey = k
		}
	}

	result := make([]int, maxCount)
	for i := maxCount - 1; i >= 0; i-- {
		result[i] = maxKey
		maxKey--
	}
	return result
}

// 最长子序列,忽略元素出现的顺序
// a:=[]int{7,1,4,3,5,9,10,25,11,12,33,2,13,6}
func MaxSubArray1(numbers []int) []int {
	if len(numbers) <= 0 {
		return numbers
	}
	m := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		currentNumber := numbers[i]
		if _, ok := m[currentNumber]; ok {
			continue
		} else if previousValue, ok := m[currentNumber-1]; ok {
			m[currentNumber] = previousValue + 1
			delete(m, currentNumber-1)
		} else {
			m[currentNumber] = 1
		}
	}

replay:
	for k, v := range m {
		if count, ok := m[k-v]; ok {
			m[k] = v + count
			delete(m, k-1)
		}
		if count, ok := m[k-v]; ok {
			m[k] = v + count
			delete(m, k-1)
			goto replay
		}
	}

	maxCount := 0
	var maxKey int
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			maxKey = k
		}
	}
	result := make([]int, maxCount)
	for i := maxCount - 1; i >= 0; i-- {
		result[i] = maxKey
		maxKey--
	}
	return result
}

// 合并区间
func BubbleSort(data [][]int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j][0] > data[j+1][0] || (data[j][0] == data[j][0] && data[j][1] > data[j+1][1]) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func Merge(ranges [][]int) [][]int {
	if len(ranges) <= 1 {
		return ranges
	}
	BubbleSort(ranges)
	result := [][]int{{ranges[0][0], ranges[0][1]}}
	for i := 1; i < len(ranges); {
		l1 := result[len(result)-1][0]
		u1 := result[len(result)-1][1]
		l2 := ranges[i][0]
		u2 := ranges[i][1]
		if u1 >= l2 {
			sub := []int{l1, int(math.Max(float64(u1), float64(u2)))}
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			result = append(result, sub)
		} else {
			result = append(result, []int{l1, u1}, []int{l2, u2})
		}
		i++
	}
	return result
}

// 最长回文子串
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
