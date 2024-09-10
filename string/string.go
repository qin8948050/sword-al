package string

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

// 字符串全排列
func Permutation1(rune []rune, begin int) {
	if begin == len(rune)-1 {
		fmt.Println(string(rune))
	}
	for i := begin; i < len(rune); i++ {
		rune[begin], rune[i] = rune[i], rune[begin]
		Permutation1(rune, begin+1)
		rune[begin], rune[i] = rune[i], rune[begin]
	}
}

func Permutation(s string) {
	if s == "" {
		return
	}
	chars := []rune(s)
	Permutation1(chars, 0)
}


// 数组中最小组合
func MinNumber(nums []int) string {
   strs:=[]string{}
   for i:=0;i<len(nums);i++ {
      strs=append(strs, strconv.Itoa(nums[i]))
   }
   sort.Slice(strs,func(i, j int) bool {
      return compare(strs[i],strs[j])
   })
   result:=strings.Join(strs,"")
   return result
}


func compare(str1 string,str2 string) bool{
   return str1+str2<str2+str1
}