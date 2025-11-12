package twosum

func twoSumTwoPointer(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right++
		}
	}
	return []int{left, right}
}

func twoSumMap(nums []int, target int) []int {
	reciprocals := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		current := target - nums[i]
		index, ok := reciprocals[current]
		if ok {
			return []int{index + 1, i + 1}
		}
		reciprocals[current] = i
	}
	return []int{}
}
