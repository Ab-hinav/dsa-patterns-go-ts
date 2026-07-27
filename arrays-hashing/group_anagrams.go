package arrayshashing

const lowercaseEnglishLetterCount = 26

// GroupAnagrams groups lowercase English words that share the same letter
// frequencies. The order of the returned groups is unspecified.
func GroupAnagrams(words []string) [][]string {
	groupsBySignature := make(map[[lowercaseEnglishLetterCount]int][]string)

	for _, word := range words {
		signature := lowercaseFrequencySignature(word)
		groupsBySignature[signature] = append(groupsBySignature[signature], word)
	}

	groups := make([][]string, 0, len(groupsBySignature))
	for _, group := range groupsBySignature {
		groups = append(groups, group)
	}

	return groups
}

func lowercaseFrequencySignature(word string) [lowercaseEnglishLetterCount]int {
	var signature [lowercaseEnglishLetterCount]int

	for _, character := range word {
		signature[character-'a']++
	}

	return signature
}
