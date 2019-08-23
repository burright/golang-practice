package bob

import (
	"fmt"
	"regexp"
)

var answerQuestion = "Sure."
var answerYell = "Whoa, chill out!"
var answerYellQuestion = "Calm down, I know what I'm doing!"
var answerBlank = "Fine. Be that way!"
var answerDefault = "Whatever."

func Hey(remark string) string {

	// Remove all whitespace and tab, newline etc. chars
	removeSpace, err := regexp.Compile("[\\s]+")

	if err != nil {
		fmt.Println(err)
	}

	remark = removeSpace.ReplaceAllString(remark, "")

	// If the remark is blank after removing whitespace answerBlank
	if len(remark) == 0 {
		return answerBlank
	}
	
	// get the trailing punctuation
	punc := string(remark[len(remark)-1])


	if punc == "?" {
		// If it is a question check if they are yelling
		if isYelling(remark) {
			return answerYellQuestion
		}
		return answerQuestion
	}

	// Check if they are yelling
	if isYelling(remark) {
		return answerYell
	}

	return answerDefault
}

func isYelling(remark string) bool {

	// Remove all non-alphabet characters 
	removeNum, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		fmt.Println(err)
	}

	remark = removeNum.ReplaceAllString(remark, "")

	// Save the number of letters before moving lower-case
	lengthBefore := len(remark)

	// Remove all lower-case letters
	removeLower, err := regexp.Compile("[^A-Z]+")
	if err != nil {
		fmt.Println(err)
	}

	remark = removeLower.ReplaceAllString(remark, "")

	// Save the number of upper-case letters
	lengthAfter := len(remark)

	// If there are no letters remaining they are not shouting
	if lengthAfter == 0 {
		return false
	}

	// If there are more or the same amount of upper-case letters
	// than lower-case + upper-case then they are yelling.
	if lengthAfter >= lengthBefore {
		return true
	}

	// Otherwise they are not yelling
	return false
}