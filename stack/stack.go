package stack

type Queue struct {
	Stack1 []int
	Stack2 []int
}

func(q *Queue)Insert(d int)  {
	q.Stack1 = append(q.Stack1,d)
}

func(q *Queue)Delete() int{
	var result int
	if len(q.Stack2)<=0{
		for len(q.Stack1)>0 {
		q.Stack2 = append(q.Stack2, q.Stack1[len(q.Stack1)-1])
		q.Stack1=q.Stack1[:len(q.Stack1)-1]
		}
	}
	if len(q.Stack2)>0{
		result=q.Stack2[len(q.Stack2)-1]
		q.Stack2=q.Stack2[:len(q.Stack2)-1]
	} else {
		panic("empty")
	}
	return result
}

type Stack struct {
	Queue1 []int
	Queue2 []int
}

func(s *Stack)Push(d int)  {
	s.Queue1 = append(s.Queue1,d)
}

func(s *Stack)Pop() int{
	if len(s.Queue1)==0  {
		panic("stack empty")
	}
	for len(s.Queue1)>1 {
		s.Queue2=append(s.Queue2, s.Queue1[0])
		s.Queue1=s.Queue1[1:]
	}
	result:=s.Queue1[0]
	s.Queue1, s.Queue2 = s.Queue2, s.Queue1[:0]
	return result
}

func VlidateStackSequences(pushed,popped []int) bool{
	if len(pushed)>0 && len(popped)>0 {
		stack:=make([]int,0)
		k:=0
		for i:=0;i<len(pushed);i++ {
			stack=append(stack,pushed[i])
			for len(stack)>0&& stack[len(stack)-1]==popped[k]{
				stack=stack[:len(stack)-1]
				k++
			}

		}
		if len(stack)==0 {
			return true
		}
	}
	return false
}