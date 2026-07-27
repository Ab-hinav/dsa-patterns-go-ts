package arrayshashing

// TwoSum returns the indices of two values whose sum equals target.
// It returns nil when no pair exists.
func TwoSum(numbers []int, target int) []int {
	seenIndices := make(map[int]int, len(numbers))

	for currentIndex, currentValue := range numbers {
		requiredValue := target - currentValue

		if matchingIndex, found := seenIndices[requiredValue]; found {
			return []int{matchingIndex, currentIndex}
		}

		seenIndices[currentValue] = currentIndex
	}

	return nil
}
