package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		expected bool
	}{
		{name: "empty input", text: "", expected: true},
		{name: "punctuation only", text: ".,,", expected: true},
		{name: "mixed case and punctuation", text: "A,b:a", expected: true},
		{name: "standard palindrome", text: "A man, a plan, a canal: Panama", expected: true},
		{name: "not a palindrome", text: "race a car", expected: false},
		{name: "digit and letter differ", text: "0P", expected: false},
		{name: "digits and mixed case", text: "1a2A1", expected: true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsPalindrome(testCase.text)

			if actual != testCase.expected {
				t.Fatalf("IsPalindrome(%q) = %t; want %t", testCase.text, actual, testCase.expected)
			}
		})
	}
}
