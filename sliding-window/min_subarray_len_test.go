package slidingwindow

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{name: "standard input", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, expected: 2},
		{name: "single qualifying element", target: 4, nums: []int{1, 4, 4}, expected: 1},
		{name: "no qualifying window", target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: 0},
		{name: "shrink after final expansion", target: 7, nums: []int{1, 1, 7}, expected: 1},
		{name: "whole array required", target: 15, nums: []int{1, 2, 3, 4, 5}, expected: 5},
		{name: "one element", target: 7, nums: []int{7}, expected: 1},
		{name: "empty input", target: 7, nums: []int{}, expected: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := MinSubArrayLen(testCase.target, testCase.nums)
			if actual != testCase.expected {
				t.Fatalf("MinSubArrayLen(%d, %v) = %d; want %d", testCase.target, testCase.nums, actual, testCase.expected)
			}
		})
	}
}
