package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(numStr string) bool {
	// Remove all spaces from the string
	removeSpace, err := regexp.Compile("[ ]+")

	if !isError(err) {
		numStr = removeSpace.ReplaceAllString(numStr, "")
	}

	// Strings of length 1 or less are not valid.
	if len(numStr) <= 1 {
		return false
	}

	// Split each number into an array of strings
	numStrArray := strings.Split(numStr, "")

	var numArray = []int{}
	// Store each string into an int array
	for _, i := range numStrArray {
		num, err := strconv.Atoi(i)
		if !isError(err) {
			numArray = append(numArray, num)
		} else {
			return false
		}
	}

	// Double each other number starting from the last number
	for i := len(numArray) - 2; i >= 0; i -= 2 {
		numArray[i] *= 2
		// If the doubled number is more than 9 then subtract 9 from it
		if numArray[i] > 9 {
			numArray[i] -= 9
		}
	}

	sum := 0
	// Add up each number in the int array
	// This could be done in the same loop where we double the numbers
	// Just need to alter the for loop to look at each index but only double every other index
	for _, val := range numArray {
		sum += val
	}

	// If the number is divisible by 10 then it is valid
	if sum % 10 == 0 {
		return true
	}

	return false
}

func isError(err error) bool {
	if err != nil {
		// If the number contains anything other than a number it is not valid
		if strings.HasSuffix(err.Error(), ": invalid syntax") {
			return true
		} 
		
		panic(err)
	}
	return false
}