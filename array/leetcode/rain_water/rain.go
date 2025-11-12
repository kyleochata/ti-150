package rainwater

/*
https://leetcode.com/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150
	in: list of elevations
	out: amt of water in the list

	2 pass
	l-r: track the heighest peak to left
	r-l: track the highest peak to right
	total water = 0
	loop
	curwater = min(pL, pR) - current height
	totalwater += max(curwater, 0)


*/

func trapWater(height []int) int {
	if len(height) < 3 {
		return 0
	}

	peakL := make([]int, len(height))
	for i := 1; i < len(peakL); i++ {
		peak := max(peakL[i-1], height[i-1])
		peakL[i] = peak
	}

	peakR := make([]int, len(height))
	for i := len(peakR) - 2; i > 0; i-- {
		peak := max(peakR[i+1], height[i+1])
		peakR[i] = peak
	}

	var totalWater int
	for i := 0; i < len(height); i++ {
		currentWater := min(peakR[i], peakL[i]) - height[i]
		totalWater += max(currentWater, 0)
	}

	return totalWater
}
