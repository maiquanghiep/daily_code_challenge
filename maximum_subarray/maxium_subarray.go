package maximumsubarray

import "math"

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

/*
As my understanding divide and conquer approach is to divide the problem to subproblems is find the maximum at current pos
To find the maximum at current pos, we compare the case of adding the current value to the sum and starting from the current value.
Repeat it to the end of the list and get the highest max value.
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxGlobal := math.MinInt32
	maxCurrent := math.MinInt32
	for _, v := range nums {
		maxCurrent = max(v, maxCurrent+v)
		maxGlobal = max(maxGlobal, maxCurrent)
	}
	return maxGlobal
}
