// https://leetcode.com/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
package main

/*
in ; slice of ints, value to remove
out: number of elements that aren't equal to val

two-pointer: start and end. Check if start is equal to value:
	a. yes: check if end is equal to val:
		aa. yes: Swap start and end. Increase start, decrease end
		ab. No: decrease end. Start main loop again
	b. No: increase the start.

	End loop when start and end indexes are equal
	Check if the start indx value is equal to the value: increase start if it doesn't match
	Return start
*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}
	end := len(nums) - 1
	start := 0
	for start < end {
		if nums[start] == val {
			if nums[end] != val {
				nums[start], nums[end] = nums[end], nums[start]
				start++
				end--
			} else {
				end--
			}
		} else {
			start++
		}
	}
	if nums[start] != val {
		start++
	}
	return start
}
