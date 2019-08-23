package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// If the sequences are not of equal length, give an error
	if (len(a) != len(b)) {
		return -1, errors.New("Sequences must be of equal length")
	}

	count := 0

	// Compare each letter and increase the count when they aren't equal
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++;
		}
	}

	// Return total count and no error
	return count, nil
}
