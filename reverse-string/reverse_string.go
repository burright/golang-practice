package reverse

import (
	"unicode/utf8"
)

func Reverse(s string) string {

	// Save the total length of runes in string
	length := utf8.RuneCountInString(s)

	// Instantiate a rune array to store reversed string
	reverse := make([]rune, length)

	i := 0
	// Iterate through the string
	for _, currRune := range s {
		// Store the current rune into the rune array
		// at position (length - i - 1)
		reverse[length - i - 1] = currRune
		i++
	}

	return string(reverse)
}
