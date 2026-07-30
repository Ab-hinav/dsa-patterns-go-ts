package twopointers

import "unicode"

// IsPalindrome reports whether text reads the same forward and backward after
// ignoring non-alphanumeric ASCII characters and letter case.
func IsPalindrome(text string) bool {
	left, right := 0, len(text)-1

	for left < right {
		if !isASCIIAlphanumeric(text[left]) {
			left++
			continue
		}

		if !isASCIIAlphanumeric(text[right]) {
			right--
			continue
		}

		leftCharacter := unicode.ToLower(rune(text[left]))
		rightCharacter := unicode.ToLower(rune(text[right]))
		if leftCharacter != rightCharacter {
			return false
		}

		left++
		right--
	}

	return true
}

func isASCIIAlphanumeric(character byte) bool {
	return (character >= 'a' && character <= 'z') ||
		(character >= 'A' && character <= 'Z') ||
		(character >= '0' && character <= '9')
}
