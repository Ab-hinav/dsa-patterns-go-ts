package twopointers

// TwoSumSorted returns the one-indexed positions of two distinct values whose
// sum equals target. Numbers must be sorted in nondecreasing order.
func TwoSumSorted(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			return []int{left + 1, right + 1}
		}

		if currentSum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
