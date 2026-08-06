package twopointers

import (
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{name: "standard input", numbers: []int{2, 7, 11, 15}, target: 9, expected: []int{1, 2}},
		{name: "outer boundaries", numbers: []int{2, 3, 4}, target: 6, expected: []int{1, 3}},
		{name: "negative and zero", numbers: []int{-1, 0}, target: -1, expected: []int{1, 2}},
		{name: "duplicate values", numbers: []int{1, 1, 3, 4}, target: 2, expected: []int{1, 2}},
		{name: "negative pair", numbers: []int{-3, -1, 0, 2, 5}, target: -4, expected: []int{1, 2}},
		{name: "does not reuse one index", numbers: []int{3, 4}, target: 6, expected: nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := TwoSumSorted(testCase.numbers, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Fatalf(
					"TwoSumSorted(%v, %d) = %v; want %v",
					testCase.numbers,
					testCase.target,
					actual,
					testCase.expected,
				)
			}
		})
	}
}
