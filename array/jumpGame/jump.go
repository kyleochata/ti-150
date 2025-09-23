package jumpgame

/*
	in: jump list
	out: boolean on if the list is jumpable

	track the furthest index you can reach
	loop through jump list:
	- if the current jump value is greater than the furthestreached index, take the jump
	- if current index is > furthestreached then we know we can't get there: return false
	edge cases:
	1. the list is empty or only one value: auto valid
	2. if the first jump can cover the whole length of the jumplist: valid
		- caution for 0 index issues (len(nums)-1)
*/

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	furthestReached := nums[0]
	if furthestReached >= len(nums)-1 {
		return true
	}
	for i, num := range nums {
		// if we reach an index > than the furthest, then return false, can't reach the current i
		if i > furthestReached {
			return false
		}
		currentFurthest := num + i
		if currentFurthest > furthestReached {
			furthestReached = currentFurthest
			if furthestReached >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
