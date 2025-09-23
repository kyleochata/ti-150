package majorityel

/*
	    in: arrray
	    out: value of majority el
		constraint: major el appears > n/2 (n == len(array))
	    naive: can stick it into a map

		Boyer-Moore Majority Vote Algorithm
		Track a single candidate at a time.
		Loop through nums
		- if val is different to the current candidate, count - 1
		- count+1 if the same val as current
		- if count < 0: track new candidate.
			- current becomes the current value the index loop is on. Reset count to 1
*/
func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
