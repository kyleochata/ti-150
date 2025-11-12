package maxarea

/*
	in: arr of heights
	out: max area
	two pointer:
	track the heights and calculate the area (use lowest height)
	save the max between old maxArea and currentArea
	move the pointer to the lowest current height
*/

func maxArea(h []int) int {
	left, right := 0, len(h)-1
	area := 0
	for left < right {
		current := min(h[left], h[right])
		width := right - left
		area = max(area, current*width)
		// Move the lowest height pointer
		if h[left] < h[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
