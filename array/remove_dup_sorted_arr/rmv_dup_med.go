package removedupsortedarr

/*
   input: ascending []int
   output: int of unique elements
   constraint: inplace; rel order maintained; O(1) extra memory; unique els can appear at max 2x
       - no extra arry must be inplace
   two pointer with a counter

   start at the 3rd el. check if two el downstream (slow-2) is different:
	a. Yes: slow become val of index fast, increase slow
	b No: increase fast

	[0,0,0,1,2,2,2,3]
	slow 2, fast 2: 0 == 0 (slow-2); fast++
	slow 2, fast 3: 1 != 0; [0, 0, 1, 1, 2, 2,2,3]; slow++
	slow 3, fast 4: 2 != 0; [0, 0, 1, 2, 2, 2, 2, 3]; slow++
	slow 4, fast 5: 2 != 1; [0, 0, 1, 2, 2, 2, 2, 3]; slow++
	slow 5, fast 6: 2 == 2 (slow-2: 3); fast++
	slow 5, fast 7: 3 != 2 (slow -2: 3 val 2); [0, 0, 1, 2, 2, 3, 2, 3]; slow++
	return 6 Nums[0,0,1,2,2,3]

*/

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
