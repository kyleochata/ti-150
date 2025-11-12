package productexecptself

/*
https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150
	in: []int
	out: []int with each i's value: product of all but self.
	constrainst: only allowed a answers array. O(1) memeory (excluding the answers arr)

	2-Pass approach
	answers array same n as nums. Set to 1
	track a running product for l->r pass. Set to 1
	L->R pass: track products of all value to left of i
		- set the answer[i] to product
			- This is the product of all values to left of i in nums
		- running product * nums[i] (this will be set in next loop)
	R -> L pass: set a running product to track products to the right of i
		- answer[i] * runningProduct
			- Mult current product (of all prod to R) to the answers[i] (this is the product of all left vals)
		- runningProduct * nums[i]
			- All products to the right of current i
*/

func productexecptself(nums []int) []int {
	if len(nums) == 0 {
		return []int{0}
	}
	answers := make([]int, len(nums))
	for i := range answers {
		answers[i] = 1
	}
	lProduct := 1
	for i := 0; i < len(answers); i++ {
		answers[i] = lProduct
		lProduct *= nums[i]
	}
	rProduct := 1
	for i := len(answers) - 1; i >= 0; i-- {
		answers[i] *= rProduct
		rProduct *= nums[i]
	}
	return answers
}
