package candy

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "empty ratings",
			ratings:  []int{},
			expected: 0,
		}, {
			name:     "single child",
			ratings:  []int{1},
			expected: 1,
		}, {
			name:     "two child asc",
			ratings:  []int{1, 2},
			expected: 3,
		},
		{
			name:     "two child asc",
			ratings:  []int{1, 2},
			expected: 3,
		},
		{
			name:     "two child des",
			ratings:  []int{2, 1},
			expected: 3,
		},
		{
			name:     "two child asc",
			ratings:  []int{1, 2},
			expected: 3,
		},
		{
			name:     "two children equal",
			ratings:  []int{1, 1},
			expected: 2, // [1, 1]
		},
		{
			name:     "three children ascending",
			ratings:  []int{1, 2, 3},
			expected: 6, // [1, 2, 3]
		},
		{
			name:     "three children descending",
			ratings:  []int{3, 2, 1},
			expected: 6, // [3, 2, 1]
		},
		{
			name:     "valley pattern",
			ratings:  []int{1, 3, 2},
			expected: 4, // [1, 2, 1]
		},
		{
			name:     "peak pattern",
			ratings:  []int{1, 2, 1},
			expected: 4, // [1, 2, 1]
		},
		{
			name:     "all equal ratings",
			ratings:  []int{2, 2, 2, 2, 2},
			expected: 5, // [1, 1, 1, 1, 1]
		},
		{
			name:     "complex case 1",
			ratings:  []int{1, 0, 2},
			expected: 5, // [2, 1, 2]
		},
		{
			name:     "complex case 2",
			ratings:  []int{1, 2, 2},
			expected: 4, // [1, 2, 1] or [1, 1, 1] but minimum is 4
		},
		{
			name:     "leetcode example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "leetcode example 2",
			ratings:  []int{1, 2, 87, 87, 87, 2, 1},
			expected: 13,
		},
		{
			name:     "strictly increasing",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15, // [1, 2, 3, 4, 5]
		},
		{
			name:     "strictly decreasing",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15, // [5, 4, 3, 2, 1]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy(tt.ratings)
			if result != tt.expected {
				t.Errorf("candy(%v) = %d, expected %d", tt.ratings, result, tt.expected)
			}
		})
	}
}

func BenchmarkCandy(b *testing.B) {
	ratings := []int{1, 2, 87, 87, 87, 2, 1, 5, 4, 3, 2, 1, 10, 8, 6, 4, 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		candy(ratings)
	}
}
