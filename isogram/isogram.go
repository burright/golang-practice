package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	// Omit dashes
	word = strings.Replace(word, "-", "", -1)
	// Omit spaces
	word = strings.Replace(word, " ", "", -1)
	// Ignore case
	word = strings.ToLower(word)

	var m = make(map[string]bool)

	// Edge case for blank word
	if word == "" { 
		return true
	}

	// Check each letter in the word
	for i := 0; i < len(word); i++ {
		// If the letter is not mapped then map it with true
		if (!m[string(word[i])]) {
			m[string(word[i])] = true
		} else {
			// Otherwise it already exists and it is not an isogram
			return false
		}
	}
	// You are an isogram!
	return true
}