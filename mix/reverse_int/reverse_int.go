package reverseint

/*

Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

*/

/*
My solution is to get sum all digit from the right, each time multiply by 10.
The optimal solution seems like to convert to string, reverse the string and convert back to int.
*/
const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if result > MaxInt32 || result < MinInt32 {
		return 0
	} else {
		return result
	}
}
