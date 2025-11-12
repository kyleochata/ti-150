package candy

/*
https://leetcode.com/problems/candy/description/?envType=study-plan-v2&envId=top-interview-150
	in: ratings; shows if the child should get more than neighbors
	out: least number of candies given

	2 PASS:
	l-r:
		number of candies the child should get more than left neighbor
		if current rating > ranting-1
			candy count for child is left child candy +1
	r-l:
		number of candies the child shoulg get more than right neighbor
		if current rating > ranting+1
			candy count for child is right child candy +1

	total up candies
*/

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	counts := make([]int, len(ratings))
	for i := range counts {
		counts[i] = 1
	}
	// start at 1; must access i-1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			counts[i] = counts[i-1] + 1
		}
	}
	// start at the end-1 (out of range if end); must access i+1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if counts[i] <= counts[i+1] {
				counts[i] = counts[i+1] + 1
			}
		}
	}
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}
