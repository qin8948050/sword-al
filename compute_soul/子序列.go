package compute_soul

// 查找数组区间，使得区间内的数字之和等于目标值
func FindSubArray(array []int, target int) (int, int, bool) {
	if len(array) == 0 {
		return -1, -1, false
	}
	start := 0
	end := 0
	currentSum := 0
	for ; end < len(array); end++ {
		currentSum += array[end]
		for currentSum > target && start <= end {
			currentSum -= array[start]
			start++
		}
		if currentSum == target {
			return start, end, true
		}
	}
	return -1, -1, false
}
