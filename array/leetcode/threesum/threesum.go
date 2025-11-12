package threesum

import (
	ms "example.com/array/sort/mergesort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	nums = ms.MergeSort(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				// Store values, not indices
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// Move both pointers after finding a valid triplet
				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func sorting(nums []int) []int {
	// Create a copy to avoid modifying original slice
	arr := make([]int, len(nums))
	copy(arr, nums)
	mergeSort(arr, 0, len(arr)-1)
	return arr
}

func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	lSize := mid - left + 1
	rSize := right - mid

	lArr := make([]int, lSize)
	rArr := make([]int, rSize)

	// Copy data to temp arrays
	for i := 0; i < lSize; i++ {
		lArr[i] = nums[left+i]
	}
	for i := 0; i < rSize; i++ {
		rArr[i] = nums[mid+1+i]
	}

	lPtr, rPtr := 0, 0
	index := left

	// Merge temp arrays back into nums
	for lPtr < lSize && rPtr < rSize {
		if lArr[lPtr] <= rArr[rPtr] {
			nums[index] = lArr[lPtr]
			lPtr++
		} else {
			nums[index] = rArr[rPtr]
			rPtr++
		}
		index++
	}

	// Copy remaining elements
	for lPtr < lSize {
		nums[index] = lArr[lPtr]
		lPtr++
		index++
	}

	for rPtr < rSize {
		nums[index] = rArr[rPtr]
		rPtr++
		index++
	}
}
