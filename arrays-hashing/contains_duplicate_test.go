package arrayshashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		expected bool
	}{
		{name: "empty input", numbers: []int{}, expected: false},
		{name: "one value", numbers: []int{1}, expected: false},
		{name: "unique values", numbers: []int{1, 2, 3, 4}, expected: false},
		{name: "adjacent duplicate", numbers: []int{1, 1, 2, 3}, expected: true},
		{name: "separated duplicate", numbers: []int{1, 2, 3, 1}, expected: true},
		{name: "negative duplicate", numbers: []int{-1, 2, -1}, expected: true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := ContainsDuplicate(testCase.numbers)

			if actual != testCase.expected {
				t.Fatalf("ContainsDuplicate(%v) = %t; want %t", testCase.numbers, actual, testCase.expected)
			}
		})
	}
}
