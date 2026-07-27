package arrayshashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "finds values at the beginning",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "uses two equal values at different indices",
			numbers:  []int{1, 2, 4, 2, 5},
			target:   4,
			expected: []int{1, 3},
		},
		{
			name:     "handles zero",
			numbers:  []int{1, 9, 0},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "returns nil when no pair exists",
			numbers:  []int{1, 2, 3},
			target:   10,
			expected: nil,
		},
		{
			name:     "returns nil for empty input",
			numbers:  []int{},
			target:   5,
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := TwoSum(testCase.numbers, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Fatalf("TwoSum(%v, %d) = %v; want %v", testCase.numbers, testCase.target, actual, testCase.expected)
			}
		})
	}
}
