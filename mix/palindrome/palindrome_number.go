package palindromell

import (
	"math"
)

/* Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Follow up: Could you solve it without converting the integer to a string?

Input: x = 121
Output: true

Input: x = -121
Output: false

My solution was to compare the first and the last digit and then remove them
from the number by subtract by the (nearest number + the last digit)/10 (For ex 1234 - 1004 = 230/10 = 23)
The mistake that I made is the case that the left digit is 0 (100021)
*/
func checkLeftRightEqual(x int, log float64) bool {
	left := x / int(math.Pow(10, math.Floor(log)))
	right := x % 10
	if left == right {
		return true
	}
	return false
}

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	log := math.Floor(math.Log10(float64(x)))

	for x >= 10 {
		left := x / int(math.Pow(10, log))
		right := x % 10

		if left != right {
			return false
		}
		x = (x - (left*int(math.Pow(10, math.Floor(log))) + right)) / 10
		log = log - 2

	}

	if x > 0 && log != 0 {
		return false
	}

	return true
}
