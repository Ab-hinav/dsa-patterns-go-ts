package arrayshashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		k        int
		expected []int
	}{
		{
			name:     "standard input",
			numbers:  []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "single value",
			numbers:  []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "negative values",
			numbers:  []int{-1, -1, 2, 2, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
		{
			name:     "all distinct values requested",
			numbers:  []int{4, 3, 2, 1},
			k:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "maximum frequency equals input length",
			numbers:  []int{7, 7, 7, 7},
			k:        1,
			expected: []int{7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := sortedIntegers(TopKFrequent(testCase.numbers, testCase.k))
			expected := sortedIntegers(testCase.expected)

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf(
					"TopKFrequent(%v, %d) = %v; want %v",
					testCase.numbers,
					testCase.k,
					actual,
					expected,
				)
			}
		})
	}
}

func sortedIntegers(numbers []int) []int {
	sorted := append([]int(nil), numbers...)
	sort.Ints(sorted)
	return sorted
}
