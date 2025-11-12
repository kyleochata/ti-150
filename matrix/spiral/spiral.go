package spiral

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	top, bot := 0, len(matrix)
	left, right := 0, len(matrix[0])
	for top < bot && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i < bot; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if !(left < right && top < bot) {
			break
		}
		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bot][i])
		}
		bot++
		for i := bot - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
