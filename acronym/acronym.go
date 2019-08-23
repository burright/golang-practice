package acronym

import (
	"regexp"
	"fmt"
	"strings"
)

func Abbreviate(s string) string {

	// Replace the hyphon from words with a space
	removeDash, err := regexp.Compile("[-]+")
	checkError(err)
	s = removeDash.ReplaceAllString(s, " ")

	// Replace the apostrophe 
	removeApos, err := regexp.Compile("[']+")
	checkError(err)
	s = removeApos.ReplaceAllString(s, "")

	// Remove everything else except letters
	onlyLetter, err := regexp.Compile("[^a-zA-Z ]+")
	checkError(err)
	s = onlyLetter.ReplaceAllString(s,"")

	// fmt.Printf("Sentence after removing punctuation: %s\n", s)

	// Split each word into an array
	words := strings.Split(s," ")

	// for i := 0; i < len(words); i++ {
	// 	fmt.Printf("\"%s\" ", string(words[i]))
	// }
	// fmt.Println("")

	acronym := ""
	word := ""

	// For each word in the array of words
	for i := 0; i < len(words); i++ {
		word = words[i]
		// If the word is not blank
		if string(words[i]) != "" {
			// Append the acronym string with the beginning of word
			acronym += string(word[0])
		}
	}

	// Capitalize all the acronym letters
	acronym = strings.ToUpper(acronym)

	return acronym
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
