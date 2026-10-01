package slidingwindow

// MinSubArrayLen returns the length of the shortest contiguous subarray whose
// sum is at least target, or 0 if none exists. target and all values in nums
// must be positive.
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	windowSum := 0
	minimumLength := len(nums) + 1

	for right, value := range nums {
		windowSum += value

		for windowSum >= target {
			minimumLength = min(minimumLength, right-left+1)
			windowSum -= nums[left]
			left++
		}
	}

	if minimumLength == len(nums)+1 {
		return 0
	}
	return minimumLength
}
