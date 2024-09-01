package analysis

import "fmt"

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
