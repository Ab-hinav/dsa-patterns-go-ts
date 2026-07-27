package arrayshashing

// ContainsDuplicate reports whether any value appears more than once.
func ContainsDuplicate(numbers []int) bool {
	seenValues := make(map[int]struct{}, len(numbers))

	for _, number := range numbers {
		if _, found := seenValues[number]; found {
			return true
		}

		seenValues[number] = struct{}{}
	}

	return false
}
