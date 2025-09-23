package rotatearr

/*
rev entire arr
rev first k el
rev n-k el
*/
import (
	"fmt"
)

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)
	reverse(nums, 0, k-1) //account for 0 index
	fmt.Println(nums)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		end--
		start++
	}
}
