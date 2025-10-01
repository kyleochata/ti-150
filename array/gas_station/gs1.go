package gasstation

func station(gas, cost []int) int {
	cTank, tTank, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		cTank += gas[i] - cost[i]
		tTank += gas[i] - cost[i]
		if cTank < 0 {
			cTank = 0
			start = i + 1
		}
	}
	if tTank < 0 {
		return -1
	}
	return start
}
