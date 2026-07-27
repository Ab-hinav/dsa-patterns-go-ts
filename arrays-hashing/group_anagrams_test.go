package arrayshashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected [][]string
	}{
		{
			name:     "empty input",
			words:    []string{},
			expected: [][]string{},
		},
		{
			name:     "empty word",
			words:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "single word",
			words:    []string{"code"},
			expected: [][]string{{"code"}},
		},
		{
			name:     "standard groups",
			words:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "preserves duplicate words",
			words:    []string{"ab", "ba", "ab"},
			expected: [][]string{{"ab", "ba", "ab"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := normalizeGroups(GroupAnagrams(testCase.words))
			expected := normalizeGroups(testCase.expected)

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("GroupAnagrams(%v) = %v; want %v", testCase.words, actual, expected)
			}
		})
	}
}

func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))

	for index, group := range groups {
		normalized[index] = append([]string(nil), group...)
		sort.Strings(normalized[index])
	}

	sort.Slice(normalized, func(leftIndex, rightIndex int) bool {
		left := normalized[leftIndex]
		right := normalized[rightIndex]

		if len(left) == 0 || len(right) == 0 {
			return len(left) < len(right)
		}

		for index := 0; index < len(left) && index < len(right); index++ {
			if left[index] != right[index] {
				return left[index] < right[index]
			}
		}

		return len(left) < len(right)
	})

	return normalized
}
