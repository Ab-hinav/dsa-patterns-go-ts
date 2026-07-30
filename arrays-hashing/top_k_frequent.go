package arrayshashing

// TopKFrequent returns the k values that occur most often.
// The order of the returned values is unspecified.
func TopKFrequent(numbers []int, k int) []int {
	frequencyByValue := make(map[int]int, len(numbers))
	for _, number := range numbers {
		frequencyByValue[number]++
	}

	valuesByFrequency := make([][]int, len(numbers)+1)
	for number, frequency := range frequencyByValue {
		valuesByFrequency[frequency] = append(valuesByFrequency[frequency], number)
	}

	result := make([]int, 0, k)
	for frequency := len(numbers); frequency > 0; frequency-- {
		for _, number := range valuesByFrequency[frequency] {
			result = append(result, number)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}
