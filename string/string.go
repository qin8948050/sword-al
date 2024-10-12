package string

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
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


//翻转单词顺序
func TraverseString(strs string) string{
   chars:=[]rune(strs)
   Reverse(chars,0,len(strs)-1)
   start:=0
   end:=0
   for end<len(chars) {
      if chars[end]==' '{
         Reverse(chars,start,end-1)
         start=end+1
      } 
      end++
   }
   Reverse(chars,start,len(chars)-1)
   return string(chars)
}

func Reverse(str []rune,start int,end int){
   for start<end {
      str[start],str[end]=str[end],str[start]
      start++
      end--
   }
}


func LeftTraverseString(strs string,input int) string{
   if input<0||input>=len(strs) {
      return strs
   }
   chars:=[]rune(strs)
   length:=len(chars)
   end:=length-1
   Reverse(chars,0,end)
   Reverse(chars,0,end-input)
   Reverse(chars,end-input+1,end)
   return string(chars)
}


// 字符串转为整数
func MyAtoi(s string) (int, error) {
	// 去掉前后的空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, errors.New("invalid input")
	}

	// 处理正负号
	sign := 1
	startIndex := 0
	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}

	result := 0
	for i := startIndex; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return 0, errors.New("invalid character encountered")
		}
		// 将字符转换为对应的数字
		digit := int(s[i] - '0')

		// 检查是否溢出（假设是32位int范围）
		if result > (math.MaxInt-digit)/10 {
			if sign == 1 {
				return math.MaxInt, nil // 正数溢出
			} else {
				return -math.MaxInt, nil // 负数溢出
			}
		}
		result = result*10 + digit
	}

	return result * sign, nil
}



// 正则匹配
func IsMatch(s string, p string) bool {
    // 如果模式为空，只有字符串也为空时才匹配
    if p == "" {
        return s == ""
    }

    // 判断首字符是否匹配（不为空且字符相同或模式是 '.'）
    firstMatch := (s != "" && (s[0] == p[0] || p[0] == '.'))

    // 如果模式的第二个字符是 '*'，那么有两种处理方法：
    if len(p) >= 2 && p[1] == '*' {
        // 1. 忽略模式中的 'x*'（即匹配 0 次），继续递归匹配 s 和 p[2:]
        // 2. 如果首字符匹配，继续递归匹配 s[1:] 和 p（即 x* 代表多个 x）
        return (IsMatch(s, p[2:]) || (firstMatch && IsMatch(s[1:], p)))
    } else {
        // 如果第二个字符不是 '*'
        // 如果首字符匹配，则继续递归匹配剩下的部分 s[1:] 和 p[1:]
        return firstMatch && IsMatch(s[1:], p[1:])
    }
}



/* 字符流中第一个出现的不重复的字符
    方法1 
 */
func firstNonRepeatingChar(stream string) string {
	m:=make(map[rune]int)
	queue:=list.New()
	for _,c:=range stream{
		m[c]++
		queue.PushBack(c)
		for queue.Len()>0 {
			front:=queue.Front()
			if m[front.Value.(rune)]>1 {
				queue.Remove(front)
			} else {
				break
			}
		}
	}
	if queue.Len()>0 {
		return string(queue.Front().Value.(rune))
	} else {
		return ""
	}
}


/* 字符流中第一个出现的不重复的字符
   方法2
 */
func firstNonRepeatingChar1(stream string) string {
	type charCount struct{
		count int
		index int
	}
	m:=make(map[rune]charCount)
	for index,s:=range stream {
		if info,found:=m[s];found {
			m[s]=charCount{count: info.count+1,index:info.index}
		} else {
			m[s]=charCount{count: 1,index:index}
		}
	}

	firstIndex:=len(stream)
	for _,info:=range m {
		if info.count==1 && >info.index<firstIndex {
			firstIndex=info.index
		}
	}
	if firstIndex<len(stream) {
		return string(stream[firstIndex])
	} else {
		return ""
	}
}