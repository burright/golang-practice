package raindrops

import "strconv"

func Convert(number int) string{
	
	// If the number is less than 3 then just return the number
	if number < 3 {
		return strconv.Itoa(number)
	}

	three := "Pling"
	five := "Plang"
	seven := "Plong"
	answer := ""

	for i := 1; i <= number; i++ {
		// Check if there is no remainder after division; therefore divisible
		if number % i == 0 {
			// If the divisible number is 3
			if i == 3 {
				answer += three
			}
			// If the divisible number is 5
			if i == 5 {
				answer += five
			}
			// If the divisible number is 7
			if i == 7 {
				answer += seven
			}
			// The requirements are to only check for numbers up to 7
			// so we can stop checking and just return what we have found
			if i > 7 { 
				if answer == "" {
					answer += strconv.Itoa(number)
				}
				
				return answer
			}
		}
	}

	return answer
}