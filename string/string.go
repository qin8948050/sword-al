package string

import "strconv"

// 千分位转换
func Parse(input int64) string {

	data := strconv.FormatInt(input, 10)
	if len(data) <= 3 {
		return data

	}
	originalIndex := len(data) - 1
	newIndex := originalIndex + (len(data) / 3)
	num := 0
	result := make([]byte, newIndex+1)
	for originalIndex >= 0 && newIndex >= originalIndex {
		if num == 3 {
			result[newIndex] = ','
			num = 0
		} else {
			result[newIndex] = data[originalIndex]
			num++
			originalIndex--
		}
		newIndex--
	}
	return string(result)
}

// 千分位转换，兼容负数
func Parse1(input int64) string {
	data := strconv.FormatInt(input, 10)

	if input >= 0 && len(data) <= 3 {
		return data
	}

	if input < 0 && len(data) <= 4 {
		return data
	}

	originalIndex := len(data) - 1
	var newIndex int
	if input < 0 {
		newIndex = originalIndex + ((len(data) - 1) / 3)
	} else {
		newIndex = originalIndex + (len(data) / 3)
	}
	num := 0
	result := make([]byte, newIndex+1)

	for originalIndex >= 0 && newIndex >= originalIndex {
		if string(data[originalIndex]) == "-" {
			result[newIndex] = '-'
			break
		}
		if num == 3 {
			result[newIndex] = ','
			num = 0
		} else {
			result[newIndex] = data[originalIndex]
			num++
			originalIndex--
		}
		newIndex--
	}

	return string(result)
}
