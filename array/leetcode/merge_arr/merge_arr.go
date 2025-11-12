package main

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	copy_n1 := make([]int, m)
	copy(copy_n1, nums1)

	left, right, curInd := 0, 0, 0
	for left < m && right < n {
		if copy_n1[left] < nums2[right] {
			left++
			curInd++
		} else {
			nums1[curInd] = nums2[right]
			right++
			curInd++
		}
	}
	for left < m {
		nums1[curInd] = copy_n1[left]
		left++
		curInd++
	}
	for right < n {
		nums1[curInd] = nums2[right]
		right++
		curInd++
	}
}
