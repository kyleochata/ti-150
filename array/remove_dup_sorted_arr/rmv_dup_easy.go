package removedupsortedarr

/*
	input: ascending sorted int
	output: number of unique el
	constraint: inplace, must preserve relative order
	slow & fast pointers:
	slow for index for swap with fast
	fast will run to try and find the next unieque value

	[0,1,1,2,3] --> [0, 1, 2, 3]
	fast = 1: 1 != 0 --> swap [0, 1, 1, 2, 3]
	fast = 2, slow = 1 --> 1 = 1 --> fast++
	fast = 3, slow 1 --> 2 != 1--> slow++ --> swap [0, 1, 2, 1, 3]
	fast = 4, slow 2 --> 3 != 2 --> slow++ --> swap [0, 1, 2, 3, 1]

	number of unique el is slow+1 (account for 0 index)

	https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
	return slow + 1
}
