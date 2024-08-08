package sort

// 数字旋转，求最小值。
func MinNum(input []int) int{
	if len(input)==0 {
		panic("input isn't empty")
	}
	head:=0
	tail:=len(input)-1
	mid:=head
	for(input[head]>=input[tail]){
		if tail-head==1 {
			mid=tail
			break
		}
		mid=(head+tail)/2
		if input[head]==input[tail] && input[head]==input[mid] {
			return minInOrder(input,head,tail)
		}
		if input[mid]>=input[head] {
			head=mid
		} else if input[mid]<=input[tail] {
			tail=mid
		}
	}
	return input[mid]
}

func minInOrder(input []int,head int,tail int) int{
	var result int
	for head<=tail {
		if input[head]<=result{
			result=input[head]
		}
		head++
	}
	return result
}