package insertintervals

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	var res [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if current[0] <= end {
			end = max(end, current[1])
		} else {
			res = append(res, []int{start, end})
			start, end = current[0], current[1]
		}
	}

	res = append(res, []int{start, end})
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
