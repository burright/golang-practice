package grains

import (
	"errors"
)

var chess = [64]uint64 {}

func Square(num int) (uint64, error) {
	if num > 64 || num <= 0 {
		return 0, errors.New("square must be between 1 and 64")
	}

	fillArray(num)

	return chess[num - 1], nil
}

func fillArray(num int) {
	for i := 0; i < num; i++ {
		if i == 0 {
			chess[i] = 1
		} else {
			chess[i] = chess[i - 1] * 2
		}
	}
}

func Total() uint64 {
	fillArray(64)
	
	var sum uint64

	for _, val := range chess {
		sum += val
	}

	return sum
}