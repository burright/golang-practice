package scrabble

import ("strconv" 
"fmt"
"strings")

// Build static scoring table
var scoring = [11][7]string {	
	{"1", "2", "3", "4", "5", "8", "10"},
	{"A", "D", "B", "F", "K", "J", "Q"},
	{"E", "G", "C", "H", "0", "X", "Z"},
	{"I", "0", "M", "V", "0", "0", "0"},
	{"O", "0", "P", "W", "0", "0", "0"},
	{"U", "0", "0", "Y", "0", "0", "0"},
	{"L", "0", "0", "0", "0", "0", "0"},
	{"N", "0", "0", "0", "0", "0", "0"},
	{"R", "0", "0", "0", "0", "0", "0"},
	{"S", "0", "0", "0", "0", "0", "0"},
	{"T", "0", "0", "0", "0", "0", "0"} }

func Score(word string) int{
	score := 0

	// Convert word to all upper to check against score table
	word = strings.ToUpper(word)

	// Check each letter's score
	for i := 0; i < len(word); i++ {
		// Keep adding up the total score
		score += Points(string(word[i]))
	}

	return score
}

func Points(letter string) int {

	// Iterate through scoring table
	for i := 1; i < 11; i++ {
		for j := 0; j < 7; j++ {
			// Check if the letter is on the score table
			if scoring[i][j] == letter {
				// If the character is in the column then return the column's score
				result, err := strconv.Atoi(scoring[0][j])

				if err != nil {
					fmt.Println("Atoi Error")
				}

				return result
			}
		}
	}

	// toDo: check for unicode characters
	return 0
}