package stack

import "unicode"

type Queue struct {
	StackIn  []int
	StackOut []int
}

func (queue *Queue) Enqueue(number int) {
	queue.StackIn = append(queue.StackIn, number)
}

func (queue *Queue) DeQueue() int {
	if len(queue.StackIn) > 0 {
		for len(queue.StackIn) > 0 {
			top := queue.StackIn[len(queue.StackIn)-1]
			queue.StackOut = append(queue.StackOut, top)
			queue.StackIn = queue.StackIn[:len(queue.StackIn)-1]
		}
	}

	if len(queue.StackOut) == 0 {
		panic("queue empty")
	}
	result := queue.StackOut[len(queue.StackOut)-1]
	queue.StackOut = queue.StackOut[:len(queue.StackOut)-1]
	return result
}

func (queue *Queue) Front() int {
	if len(queue.StackIn) > 0 {
		for len(queue.StackIn) > 0 {
			top := queue.StackIn[len(queue.StackIn)-1]
			queue.StackOut = append(queue.StackOut, top)
			queue.StackIn = queue.StackIn[:len(queue.StackIn)-1]
		}
	}

	if len(queue.StackOut) == 0 {
		panic("queue empty")
	}
	result := queue.StackOut[len(queue.StackOut)-1]
	return result
}

type Stack struct {
	Queue1 []int
	Queue2 []int
}

func (s *Stack) Push(d int) {
	s.Queue1 = append(s.Queue1, d)
}

func (s *Stack) Pop() int {
	if len(s.Queue1) == 0 {
		panic("stack empty")
	}
	for len(s.Queue1) > 1 {
		s.Queue2 = append(s.Queue2, s.Queue1[0])
		s.Queue1 = s.Queue1[1:]
	}
	result := s.Queue1[0]
	s.Queue1, s.Queue2 = s.Queue2, s.Queue1[:0]
	return result
}

func (s *Stack) Top() int {
	if len(s.Queue1) == 0 {
		panic("stack empty")
	}
	var result int
	for len(s.Queue1) > 0 {
		result = s.Queue1[0]
		s.Queue2 = append(s.Queue2, s.Queue1[0])
		s.Queue1 = s.Queue1[1:]
	}
	s.Queue1, s.Queue2 = s.Queue2, s.Queue1
	return result
}

func ValidateStackSequences(pushed, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	stack := make([]int, 0)
	k := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[k] {
			stack = stack[:len(stack)-1]
			k++
		}

	}
	return len(stack) == 0
}

// 基本计算器
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
