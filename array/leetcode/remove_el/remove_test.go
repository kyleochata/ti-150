package main

import (
	// "reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		wantNums []int // Expected nums slice up to the returned length
	}{
		{
			name:     "empty array",
			nums:     []int{},
			val:      3,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "single element matching value",
			nums:     []int{3},
			val:      3,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "single element not matching value",
			nums:     []int{3},
			val:      2,
			expected: 1,
			wantNums: []int{3},
		},
		{
			name:     "all elements match value",
			nums:     []int{2, 2, 2, 2},
			val:      2,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "no elements match value",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			expected: 4,
			wantNums: []int{1, 2, 3, 4},
		},
		{
			name:     "some elements match value at beginning",
			nums:     []int{3, 3, 1, 2},
			val:      3,
			expected: 2,
			wantNums: []int{1, 2},
		},
		{
			name:     "some elements match value at end",
			nums:     []int{1, 2, 3, 3},
			val:      3,
			expected: 2,
			wantNums: []int{1, 2},
		},
		{
			name:     "some elements match value in middle",
			nums:     []int{1, 3, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{1, 2},
		},
		{
			name:     "mixed elements with duplicates",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			wantNums: []int{0, 1, 4, 0, 3},
		},
		{
			name:     "leetcode example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{2, 2},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			wantNums: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input slice to avoid modifying the test case
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := removeElement(numsCopy, tt.val)

			if got != tt.expected {
				t.Errorf("removeElement() = %v, want %v", got, tt.expected)
			}

			// Check that the first 'got' elements match the expected values
			// Note: The order of non-removed elements might vary, but they should be present
			if got > 0 {
				// For a more robust test, we could check if all expected values are present
				// in the result, regardless of order
				if !containsAll(numsCopy[:got], tt.wantNums) {
					t.Errorf("removeElement() modified slice = %v, want first %v elements to contain %v",
						numsCopy[:got], got, tt.wantNums)
				}
			} else if got == 0 && len(tt.wantNums) > 0 {
				t.Errorf("removeElement() returned 0 but expected non-empty result: %v", tt.wantNums)
			}

			// Additional check: verify that no removed values exist in the first 'got' elements
			for i := 0; i < got; i++ {
				if numsCopy[i] == tt.val {
					t.Errorf("removeElement() left value %v at index %d in result slice %v",
						tt.val, i, numsCopy[:got])
				}
			}
		})
	}
}

// Helper function to check if slice a contains all elements of slice b
func containsAll(a, b []int) bool {
	if len(a) < len(b) {
		return false
	}

	// Create frequency maps
	countA := make(map[int]int)
	countB := make(map[int]int)

	for _, num := range a {
		countA[num]++
	}
	for _, num := range b {
		countB[num]++
	}

	// Check if a has at least as many of each element as b
	for num, count := range countB {
		if countA[num] < count {
			return false
		}
	}
	return true
}

// Edge case tests
func TestRemoveElementEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "large numbers",
			nums:     []int{1000000, 999999, 1000000},
			val:      1000000,
			expected: 1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -1},
			val:      -1,
			expected: 2,
		},
		{
			name:     "zero value",
			nums:     []int{0, 1, 0, 3, 0},
			val:      0,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := removeElement(numsCopy, tt.val)

			if got != tt.expected {
				t.Errorf("removeElement() = %v, want %v", got, tt.expected)
			}

			// Verify no instances of val in the result
			for i := 0; i < got; i++ {
				if numsCopy[i] == tt.val {
					t.Errorf("Found value %v at index %d in result", tt.val, i)
				}
			}
		})
	}
}

// Benchmark test
func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 2, 1, 5, 6, 2, 7, 8, 2}
	val := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		removeElement(numsCopy, val)
	}
}
