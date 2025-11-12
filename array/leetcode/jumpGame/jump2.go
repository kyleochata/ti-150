package jumpgame

/*
https://leetcode.com/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150
	in: array of jump values
	out: least number of jumps
	constraint: guranteed to reach n-1

	furthest : set to nums[0]
		Window to show the range of indexes upcoming to check for the nextFurthest jump
	nextFurthest: set to nums[0]
		largest jump value from current i to furthest
	loop starting at i := 1:
		check if current index + value is > nextFurthest --> nextFurthest = current
		when i == furthest: at the end of jump window, must take a jump
			count++
			furthest = nextfurthest
		check if new futhest >= len(n)-1

*/

func jump2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	furthest := nums[0]
	if furthest >= len(nums)-1 {
		return 1
	}

	nextFurthest := nums[0]
	var count int

	for i := 1; i < len(nums); i++ {
		current := nums[i] + i
		if current > nextFurthest {
			nextFurthest = current
		}
		if i == furthest {
			furthest = nextFurthest
			count++
		}
		if furthest >= len(nums)-1 {
			return count
		}
	}
	return count
}
