package stack

type Queue struct {
   StackIn []int
   StackOut []int
}

func(queue *Queue)Enqueue(number int) {
   queue.StackIn=append(queue.StackIn, number)
}

func(queue *Queue)DeQueue() int {
   if len(queue.StackIn)>0{
      for len(queue.StackIn)>0 {
         top:=queue.StackIn[len(queue.StackIn)-1]
         queue.StackOut=append(queue.StackOut, top)
         queue.StackIn=queue.StackIn[:len(queue.StackIn)-1]
      }
   }
   
   if len(queue.StackOut)==0 {
      panic("queue empty")
   } 
   result:=queue.StackOut[len(queue.StackOut)-1]
   queue.StackOut=queue.StackOut[:len(queue.StackOut)-1]
   return result
}

func(queue *Queue)Front() int {
   if len(queue.StackIn)>0{
      for len(queue.StackIn)>0 {
         top:=queue.StackIn[len(queue.StackIn)-1]
         queue.StackOut=append(queue.StackOut, top)
         queue.StackIn=queue.StackIn[:len(queue.StackIn)-1]
      }
   }
   
   if len(queue.StackOut)==0 {
      panic("queue empty")
   } 
   result:=queue.StackOut[len(queue.StackOut)-1]
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

func(s *Stack)Top() int {
	if len(s.Queue1)==0  {
		panic("stack empty")
	}
	var result int
	for len(s.Queue1)>0 {
		result=s.Queue1[0]
		s.Queue2=append(s.Queue2, s.Queue1[0])
		s.Queue1=s.Queue1[1:]
	}
	s.Queue1, s.Queue2 = s.Queue2, s.Queue1
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