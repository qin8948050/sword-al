package main

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	stack := []int{}
	sign := '+'
	num := 0

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = int(s[i] - '0')
			for i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
				i++
				num = num*10 + int(s[i]-'0')
			}
		}

		if (!unicode.IsDigit(rune(s[i])) && s[i] != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			sign = rune(s[i])
			num = 0
		}

	}
	result := 0
	for _, value := range stack {
		result += value
	}
	return result
}

func main() {
	result := calculate("42")
	fmt.Println(result)
}
