package isogram

import (
	"strings"
)

var m = make(map[string]int)

func IsIsogram(word string) bool {
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)

	if word == "" { 
		return true
	}

	for i := 0; i < len(word); i++ {
		if val, ok := m[string(word[i])]; !ok {
			m[string(val)] = 1
		} else {
			return false
		}

		return true
	}
	return false
}