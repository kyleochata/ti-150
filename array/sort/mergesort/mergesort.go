package mergesort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	mergesort(sorted, 0, len(sorted)-1)
	return sorted
}

func mergesort(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergesort(nums, left, mid)
		mergesort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	lsize := mid - left + 1
	rsize := right - mid
	lArr := make([]int, lsize)
	rArr := make([]int, rsize)
	// Copy data to temporary arrays
	for i := 0; i < lsize; i++ {
		lArr[i] = nums[left+i]
	}
	for i := 0; i < rsize; i++ {
		rArr[i] = nums[mid+1+i]
	}

	// Merge the temporary arrays back
	i, j, k := 0, 0, left
	for i < lsize && j < rsize {
		if lArr[i] <= rArr[j] {
			nums[k] = lArr[i]
			i++
		} else {
			nums[k] = rArr[j]
			j++
		}
		k++
	}

	// Copy remaining elements from lArr
	for i < lsize {
		nums[k] = lArr[i]
		i++
		k++
	}

	// Copy remaining elements from rArr
	for j < rsize {
		nums[k] = rArr[j]
		j++
		k++
	}

}
