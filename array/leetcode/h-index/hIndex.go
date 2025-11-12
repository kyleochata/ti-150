package hindex

/*
https://leetcode.com/problems/h-index/submissions/1780648562/?envType=study-plan-v2&envId=top-interview-150
	in: array with citations. i = paper number, val = citation count for that paper
	out: h-index
	h-index: Largest number where there's at least h papers with at least h citations

	create buckets of n+1 size.
		- Papers go into  buckets based on citation #
		- If citation # > n then add to last bucket
	Reverse accumulation:
		- start at end of bucket array.
		- add citation count to total papers.
		- index of bucket arry is the number of citations
		- if the number of citations is less than or equal to the number of papers -> h-index

*/

func hIndex(citations []int) int {
	length := len(citations)
	if length == 0 {
		return 0
	}
	buckets := make([]int, length+1)
	for _, citation := range citations {
		if citation >= length {
			buckets[length] += 1
		} else {
			buckets[citation] += 1
		}
	}
	totalPapers := 0
	for i := length; i >= 0; i-- {
		totalPapers += buckets[i]
		if totalPapers >= i {
			return i
		}
	}
	return 0
}
