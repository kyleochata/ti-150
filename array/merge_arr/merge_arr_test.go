package main

import (
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		m        int
		n        int
		expected []int
	}{
		{
			name:     "simple sort",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{4, 5, 6},
			m:        3,
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
